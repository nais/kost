package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Health check ping")
	respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func CostHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Fetching mock cost")
	respondJSON(w, http.StatusOK, map[string]float64{"gcp_cost": 42.00})
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
