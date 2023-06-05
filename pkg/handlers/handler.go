package handlers

import (
	"encoding/json"
	"strconv"
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
}

func New(db *gorm.DB, slackWebhookURL, slackChannel string) *handler {
	return &handler{
		DB:              db,
		slackWebhookURL: slackWebhookURL,
		slackChannel:    slackChannel,
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
			text, err := json.MarshalIndent(alert, "  ", "  ")
			if err != nil {
				log.Error("failed to marshal Alert", "err", err)
				continue
			}
			attachment := slack.Attachment{
				Color:      "good",
				Fallback:   "Receive Forta Alerts",
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
