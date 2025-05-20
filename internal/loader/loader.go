package loader

import (
	"encoding/csv"
	"log"
	"os"
	"sales-api/internal/db"
	"strconv"
	"time"
)

func RefreshData(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Read() // skip header

	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// overwriting existing data or appending new data while managing duplicates
	stmtCust, _ := tx.Prepare(`INSERT OR REPLACE INTO customers VALUES (?, ?, ?, ?, ?)`)
	stmtProd, _ := tx.Prepare(`INSERT OR REPLACE INTO products VALUES (?, ?, ?)`)
	stmtOrder, _ := tx.Prepare(`INSERT OR REPLACE INTO orders VALUES (?, ?, ?, ?, ?)`)
	stmtItem, _ := tx.Prepare(`INSERT INTO order_items (order_id, product_id, quantity, unit_price, discount) VALUES (?, ?, ?, ?, ?)`)

	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		orderID, _ := strconv.Atoi(record[0])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shipping, _ := strconv.ParseFloat(record[10], 64)
		quantity, _ := strconv.Atoi(record[7])
		orderDate, _ := time.Parse("2006-01-02", record[6])

		stmtCust.Exec(record[2], record[12], record[13], record[14], record[5])
		stmtProd.Exec(record[1], record[3], record[4])
		stmtOrder.Exec(orderID, record[2], orderDate.Format("2006-01-02"), record[11], shipping)
		stmtItem.Exec(orderID, record[1], quantity, unitPrice, discount)
	}
	tx.Commit()
}