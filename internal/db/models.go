package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./sales.db")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := DB.Exec(schema); err != nil {
		log.Fatal(err)
	}
}

const schema = `
CREATE TABLE IF NOT EXISTS customers (
	id TEXT PRIMARY KEY,
	name TEXT,
	email TEXT,
	address TEXT,
	region TEXT
);
CREATE TABLE IF NOT EXISTS products (
	id TEXT PRIMARY KEY,
	name TEXT,
	category TEXT
);
CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY,
	customer_id TEXT,
	order_date DATE,
	payment_method TEXT,
	shipping_cost REAL
);
CREATE TABLE IF NOT EXISTS order_items (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	order_id INTEGER,
	product_id TEXT,
	quantity INTEGER,
	unit_price REAL,
	discount REAL
);`