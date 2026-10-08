# Clinic Beauty Management — Backend

Go modular monolith for the clinic/beauty product.

**Status:** Auth module working with demo users.  
No production money/stock logic yet. Accounting path is in `docs/ACCOUNTING-PATH.md`.

## Requirements

- Go 1.22+
- Docker + Docker Compose

## Quick start (local)

```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

API listens on `http://localhost:8080`.

## Auth (ready for testing)

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/v1/auth/login` | No | Login, returns JWT |
| GET | `/api/v1/auth/me` | Bearer | Current user info |
| GET | `/api/v1/auth/ping` | No | Smoke test |

### Demo accounts (local only)

| Username | Password | Role |
|----------|----------|------|
| manager | manager123 | manager |
| reception | reception123 | reception |
| doctor | doctor123 | doctor |
| cashier | cashier123 | cashier |

Example:

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"manager","password":"manager123"}'
```

Then use the token:

```bash
curl http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer <access_token>"
```

## Architecture principles

- **Modular Monolith** — each feature lives in `internal/modules/<name>`
- Modules communicate only through interfaces
- Adding a new UI card + API endpoint = change only inside one module
- Single clinic only

## Accounting hard rules

1. Opening stock ≠ purchase
2. Confirmed deposit → patient wallet credit (liability), not revenue
3. Soft-void only + audit who/when
4. Stock valuation: Weighted Average

See `docs/ACCOUNTING-PATH.md`.

## Next steps

1. Patients CRUD + wallet read model
2. Admissions
3. After accounting checklist Done → finance & inventory
