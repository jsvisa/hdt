package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/jsvisa/hdt/pkg/models"
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
	// Append to the Books table
	if result := h.DB.CreateInBatches(&rpcAlert.Alerts, 10); result.Error != nil {
		log.Error("failed to save alerts into db", "err", result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
