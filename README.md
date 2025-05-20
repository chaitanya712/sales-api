# Sales Data API in Go

## How to Run
```bash
go mod tidy
go run cmd/main.go
```

## API
| Route | Method | Params | Description |
|-------|--------|--------|-------------|
| /revenue | GET | start, end (YYYY-MM-DD) | Total revenue between dates |
| /revenue/by-product | GET | start, end (YYYY-MM-DD) | Total revenue grouped by product |
| /revenue/by-category | GET | start, end (YYYY-MM-DD) | Total revenue grouped by category |
| /revenue/by-region | GET | start, end (YYYY-MM-DD) | Total revenue grouped by region |


## Notes
- DB file: `sales.db`
- Input CSV: `data/sales.csv`
- Ensure `data/sales.csv` exists with valid headers
