package main

import (
	"log"
	"net/http"

	"sales-api/internal/api"
	"sales-api/internal/db"
	"sales-api/internal/loader"
)

func main() {
	db.InitDB()
	loader.RefreshData("data/sales.csv")
	http.HandleFunc("/revenue", api.RevenueHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
