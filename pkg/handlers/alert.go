package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jsvisa/hdt/pkg/models"
	"gorm.io/datatypes"
)

var (
	// Ref https://github.com/DefiLlama/chainlist/blob/main/constants/chainIds.json
	CHAINIDS = map[uint64]string{
		1:     "ethererum",
		10:    "optimistic",
		25:    "cronos",
		56:    "bsc",
		66:    "okex",
		128:   "heco",
		137:   "bor",
		250:   "fantom",
		42161: "arbitrum",
		42220: "celo",
		42170: "nova",
		43114: "avalanche",
	}
)

func (h handler) AddAlert(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("failed to read body", "err", err)
		return
	}

	var rpcAlert models.RPCAlerts
	err = json.Unmarshal(body, &rpcAlert)
	if err != nil {
		log.Error("failed to unmarshal", "err", err)
	}

	log.Info("recv alerts", "#alert", len(rpcAlert.Alerts))
	if len(rpcAlert.Alerts) > 0 {
		alerts := make([]*models.Alert, len(rpcAlert.Alerts))
		for i, alert := range rpcAlert.Alerts {
			alerts[i] = &models.Alert{
				AlertID:     alert.AlertID,
				Name:        alert.Name,
				Description: alert.Description,
				CreatedAt:   alert.CreatedAt,
				FindingType: alert.FindingType,
				Severity:    alert.Severity,
				Metadata:    datatypes.JSON(alert.Metadata),
			}
			if source := alert.Source; source != nil {
				alerts[i].TxHash = source.TransactionHash
				if block := source.Block; block != nil {
					chainID := block.ChainID
					if chain, ok := CHAINIDS[chainID]; ok {
						alerts[i].Chain = chain
					}
					alerts[i].BlockNum = block.Number
					if blockTime, err := hexutil.DecodeUint64(block.Timestamp); err == nil {
						alerts[i].BlockTimestamp = time.Unix(int64(blockTime), 0)
					} else {
						log.Error("failed to decode block timestamp", "block.Timestamp", block.Timestamp, "err", err)
						alerts[i].BlockTimestamp = time.Now()
					}
				}
			}
		}
		h.postAlerts(alerts)

		// Append to the alerts table
		if result := h.DB.CreateInBatches(&alerts, 10); result.Error != nil {
			log.Error("failed to save alerts into db", "err", result.Error)
		}
	}

	// Send a 200 OK
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Created")
}
