package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/slack-go/slack"
	"gorm.io/gorm"

	"github.com/jsvisa/hdt/pkg/models"
)

type handler struct {
	DB              *gorm.DB
	slackWebhookURL string
	slackChannel    string
	slackSeverities map[string]bool
}

func New(db *gorm.DB, slackWebhookURL, slackChannel, slackSeverity string) *handler {
	severities := make(map[string]bool)
	for _, s := range strings.Split(slackSeverity, ",") {
		severities[strings.ToUpper(s)] = true
	}
	return &handler{
		DB:              db,
		slackWebhookURL: slackWebhookURL,
		slackChannel:    slackChannel,
		slackSeverities: severities,
	}
}

func (h *handler) slackEnabled() bool {
	return len(h.slackWebhookURL) > 0 && len(h.slackChannel) > 0
}

func (h *handler) postAlerts(alerts []*models.Alert) {
	if !h.slackEnabled() {
		return
	}
	go func() {
		for _, alert := range alerts {
			if _, ok := h.slackSeverities[alert.Severity]; !ok {
				continue
			}
			a := map[string]interface{}{
				"name":            alert.Name,
				"chain":           alert.Chain,
				"block_timestamp": alert.BlockTimestamp,
				"blknum":          alert.BlockNum,
				"txhash":          alert.TxHash,
				"findingType":     alert.FindingType,
				"severity":        alert.Severity,
				"description":     alert.Description,
				"metadata":        alert.Metadata,
			}
			text, err := json.MarshalIndent(a, "  ", "  ")
			if err != nil {
				log.Error("failed to marshal Alert", "err", err)
				continue
			}
			attachment := slack.Attachment{
				Color:      "good",
				Fallback:   fmt.Sprintf("Receive Forta Alert: %s", alert.AlertID),
				Text:       string(text),
				Footer:     "forta alert",
				FooterIcon: ":microbe:",
				Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
			}
			msg := slack.WebhookMessage{
				Channel:     h.slackChannel,
				Attachments: []slack.Attachment{attachment},
			}

			if err := slack.PostWebhook(h.slackWebhookURL, &msg); err != nil {
				log.Error("post slack error", "err", err)
				if retryErr, ok := err.(*slack.RateLimitedError); ok {
					time.Sleep(retryErr.RetryAfter)
				}
			}
		}
	}()
}
