package handlers

import (
	"encoding/json"
	"github.com/stickpro/vending/pkg/logger"
	"net/http"
)

func UsersPageIndex(w http.ResponseWriter, r *http.Request) {
	users, err :=
	if err != nil {
		logger.Error(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
