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

## Notes
- DB file: `sales.db`
- Input CSV: `data/sales.csv`
- Ensure `data/sales.csv` exists with valid headers
