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

example of Request url
```bash
http://localhost:8080/revenue?start=2023-12-01&end=2024-06-01
```


## Notes
- DB file: `sales.db`
- Input CSV: `data/sales.csv`
- Ensure `data/sales.csv` exists with valid headers
