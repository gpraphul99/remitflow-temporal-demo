# remitflow-temporal-demo

A tiny but realistic cross-border remittance flow built with Go + Temporal just for fun (and to learn their SDK properly).

Turns ¥ → USD (or whatever), does KYC, compliance, payout, and automatically refunds if anything fails — exactly how real payment systems should work.

### Stack
- Go 1.23
- Temporal Go SDK
- Chi router
- PostgreSQL (pgx)
- docker-compose (Temporal server + DB + Web UI included)

### Quick start
```bash
docker compose up -d          # spins up Temporal + Postgres + Web UI
go run worker/main.go         # terminal 1
go run api/main.go            # terminal 2 → listens on :8080