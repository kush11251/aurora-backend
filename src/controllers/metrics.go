package controllers

import (
	"aurora-backend/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := models.GetFinancialMetrics()
	json.NewEncoder(w).Encode(metrics)
}