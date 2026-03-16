package main

import (
	"aurora-backend/src/controllers"
	"aurora-backend/src/models"
	"aurora-backend/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/metrics", controllers.GetMetrics)
	http.ListenAndServe(":8080", nil)
}