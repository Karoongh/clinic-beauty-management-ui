# Clinic Beauty Management — Backend

Go modular monolith for the clinic/beauty product.

**Status:** Skeleton + non-money module stubs.  
No production money/stock logic yet. Accounting path is documented in `docs/ACCOUNTING-PATH.md`.

## Requirements

- Go 1.22+
- Docker + Docker Compose

## Quick start (local)

```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

API listens on `http://localhost:8080`.

| Endpoint | Purpose |
|----------|---------|
| `GET /health` | Health check |
| `GET /api/v1/auth/ping` | Auth module alive |
| `GET /api/v1/patients/ping` | Patients module alive |
| `GET /api/v1/admissions/ping` | Admissions module alive |

## Architecture principles

- **Modular Monolith** — each feature lives in `internal/modules/<name>`
- Modules communicate only through interfaces (no direct imports between modules)
- Adding a new UI card + API endpoint = change only inside one module
- Stateless API + Redis + PostgreSQL connection pool for high concurrency
- Single clinic only (no multi-tenant ERP)

## Accounting hard rules (do not violate)

1. Opening stock ≠ purchase
2. Confirmed deposit → patient wallet credit (liability), not service revenue
3. All money/stock documents: unique sequence, soft-void, audit who/when
4. Stock valuation: Weighted Average

See `docs/ACCOUNTING-PATH.md` and root `docs/ACCOUNTING.md`.

## Project layout

```
backend/
├── cmd/api/                 # entrypoint
├── internal/
│   ├── config/
│   ├── platform/            # db, redis
│   ├── modules/             # feature modules (isolated)
│   │   ├── auth/
│   │   ├── patients/
│   │   ├── admissions/
│   │   └── finance/         # structure only – gated
│   └── shared/
├── migrations/
├── deploy/
│   ├── docker/
│   └── compose/
└── docs/
```

## Next steps

1. Implement real auth (login, JWT, roles)
2. Implement patients CRUD + wallet read model
3. Implement admissions
4. Only after accounting checklist is fully Done → finance & inventory code
