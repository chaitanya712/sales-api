package api

import (
	"fmt"
	"net/http"
	"sales-api/internal/db"
	
)

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
