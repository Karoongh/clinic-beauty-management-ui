# Clinic Beauty Management — Backend

Go modular monolith. Demo data. **No production money/stock mutation.**

## Quick start
```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

## Ready modules (all under `/api/v1`)

| Module | Key endpoints |
|--------|----------------|
| Auth | `POST /auth/login` · `GET /auth/me` |
| Patients | `GET /patients` · `GET /patients/:id` · `GET /patients/:id/wallet` |
| Admissions | `GET /admissions` · `GET /admissions/today` · `POST /admissions` |
| Appointments | `GET /appointments/requests` · `POST /appointments/requests/:id/decide` |
| Catalog | `GET /catalog/items` · `GET /catalog/items/:id` |
| Analytics | `GET /analytics/home` |

Demo logins: `manager/manager123` · `reception/reception123` · `doctor/doctor123` · `cashier/cashier123`

## Accounting gate (still open)

Finance and inventory **must not** be implemented until every item in `docs/ACCOUNTING-PATH.md` is marked Done and reviewed.

Hard rules remain:
- Opening stock ≠ purchase
- Deposit → wallet credit (liability), not revenue
- Soft-void + audit who/when
- Weighted Average valuation

## Safety

Work lives on branch `feat/backend-skeleton-go`.  
Safety snapshot: `safety/20261008-backend-skeleton`
