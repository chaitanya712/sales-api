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
	http.HandleFunc("/revenue/by-product", api.RevenueByProductHandler)
	http.HandleFunc("/revenue/by-category", api.RevenueByCategoryHandler)
	http.HandleFunc("/revenue/by-region", api.RevenueByRegionHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
