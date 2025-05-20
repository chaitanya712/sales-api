package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sales-api/internal/db"
)

type RevenueGroup struct {
	Key    string  `json:"key"`
	Amount float64 `json:"amount"`
}

func RevenueHandler(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	if start == "" || end == "" {
		http.Error(w, "start and end date required", http.StatusBadRequest)
		return
	}
	q := `SELECT quantity, unit_price, discount FROM order_items 
		JOIN orders ON order_items.order_id = orders.id 
		WHERE order_date BETWEEN ? AND ?`
	rows, err := db.DB.Query(q, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	total := 0.0
	for rows.Next() {
		var qty int
		var price, disc float64
		rows.Scan(&qty, &price, &disc)
		total += float64(qty) * price * (1 - disc)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"total_revenue": %.2f}`, total)
}

func RevenueByProductHandler(w http.ResponseWriter, r *http.Request) {
	generateGroupedRevenue(w, r, "products.name")
}

func RevenueByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	generateGroupedRevenue(w, r, "products.category")
}

func RevenueByRegionHandler(w http.ResponseWriter, r *http.Request) {
	generateGroupedRevenue(w, r, "customers.region")
}

func generateGroupedRevenue(w http.ResponseWriter, r *http.Request, groupBy string) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	if start == "" || end == "" {
		http.Error(w, "start and end date required", http.StatusBadRequest)
		return
	}

	q := fmt.Sprintf(`SELECT %s, SUM(quantity * unit_price * (1 - discount)) as revenue
		FROM order_items
		JOIN orders ON order_items.order_id = orders.id
		JOIN products ON order_items.product_id = products.id
		JOIN customers ON orders.customer_id = customers.id
		WHERE order_date BETWEEN ? AND ?
		GROUP BY %s`, groupBy, groupBy)

	rows, err := db.DB.Query(q, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []RevenueGroup
	for rows.Next() {
		var group RevenueGroup
		rows.Scan(&group.Key, &group.Amount)
		results = append(results, group)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}