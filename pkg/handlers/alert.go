package handlers

import (
	"encoding/json"
	"fmt"
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

	var alert models.Alert
	json.Unmarshal(body, &alert)

	// Append to the Books table
	if result := h.DB.Create(&alert); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
