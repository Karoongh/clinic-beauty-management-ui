# Clinic Beauty Management — Backend

Go modular monolith for the clinic/beauty product.

**Status:** Auth + Patients modules ready (demo data).  
No production money/stock mutation yet.

## Requirements

- Go 1.22+
- Docker + Docker Compose

## Quick start (local)

```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

API listens on `http://localhost:8080`.

## Auth

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/v1/auth/login` | No | Login → JWT |
| GET | `/api/v1/auth/me` | Bearer | Current user |
| GET | `/api/v1/auth/ping` | No | Smoke test |

### Demo accounts

| Username | Password | Role |
|----------|----------|------|
| manager | manager123 | manager |
| reception | reception123 | reception |
| doctor | doctor123 | doctor |
| cashier | cashier123 | cashier |

## Patients

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/patients?q=` | Bearer | List / search |
| GET | `/api/v1/patients/:id` | Bearer | Profile |
| GET | `/api/v1/patients/:id/wallet` | Bearer | Wallet balance (read-only) |
| GET | `/api/v1/patients/ping` | No | Smoke test |

Example:

```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"manager","password":"manager123"}' | jq -r .access_token)

curl http://localhost:8080/api/v1/patients \
  -H "Authorization: Bearer $TOKEN"

curl "http://localhost:8080/api/v1/patients?q=سارا" \
  -H "Authorization: Bearer $TOKEN"

curl http://localhost:8080/api/v1/patients/p1/wallet \
  -H "Authorization: Bearer $TOKEN"
```

## Architecture principles

- Modular Monolith – each feature in `internal/modules/<name>`
- Modules talk only through interfaces
- Adding a UI card + endpoint = change only inside one module
- Single clinic only

## Accounting hard rules

1. Opening stock ≠ purchase
2. Deposit → wallet credit (liability), not revenue
3. Soft-void + audit who/when
4. Stock valuation: Weighted Average

See `docs/ACCOUNTING-PATH.md`.

## Next steps

1. Admissions module
2. After accounting checklist Done → finance & inventory
