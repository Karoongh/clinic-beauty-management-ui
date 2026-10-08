# Clinic Beauty Management — Backend

Go modular monolith for the clinic/beauty product.

**Status:** Skeleton only. No production money/stock logic yet.
Accounting rules are documented and must pass `accounting-audit-dev` gates before any finance code is written.

## Requirements

- Go 1.22+
- Docker + Docker Compose

## Quick start (local)

```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

API will listen on `http://localhost:8080`.

Health check: `GET /health`

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
4. Stock valuation: Weighted Average (documented)

See `docs/ACCOUNTING.md` in the repository root and `internal/modules/finance/README.md`.

## Project layout

```
backend/
├── cmd/api/                 # entrypoint
├── internal/
│   ├── config/
│   ├── platform/            # db, redis, logger, auth middleware
│   ├── modules/             # feature modules (isolated)
│   └── shared/
├── migrations/
├── deploy/
│   ├── docker/
│   └── compose/
└── api/
```

## Next steps after skeleton

1. Complete accounting path + checklist (accounting-audit-dev)
2. Implement auth + patients + admissions
3. Then finance/inventory under full audit gates
