# Thanks for considering me in your job applicant interviews. 🤍🙏


## Setup 
copy `env.example` to `.env` and set your postgreSQL connection data <br><br>
apply migration on your database 
```go
make db-migrate-up
```
run 
```go
go run ./cmd/main.go
```
## Check out wit curl 
```go
curl -X POST -H "Content-Type: application/json" -d '{"Id":1,
  "instrumentId": 1,
  "dateEn": "2023-10-31T12:00:00Z",
  "open": 10.5,
  "high": 11.2,
  "low": 10.1,
  "close": 10.8
}' http://localhost:8000/api/trade
```

# Run unit test 
```go
make unit_test
```