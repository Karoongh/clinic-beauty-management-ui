# Clinic Beauty Management — Backend

Go modular monolith. Demo data. No production money mutation yet.

## Quick start
```bash
cd backend && docker compose -f deploy/compose/docker-compose.yml up --build
```

## Modules ready

| Module | Endpoints |
|--------|-----------|
| **Auth** | `POST /auth/login` · `GET /auth/me` |
| **Patients** | `GET /patients?q=` · `GET /patients/:id` · `GET /patients/:id/wallet` |
| **Admissions** | `GET /admissions` · `GET /admissions/today` · `POST /admissions` |
| **Appointments** | `GET /appointments/requests` · `POST /appointments/requests/:id/decide` |
| **Catalog** | `GET /catalog/items?kind=&q=` · `GET /catalog/items/:id` |

All under `/api/v1` and protected by Bearer token (except login & pings).

Demo logins: `manager/manager123`, `reception/reception123`, `doctor/doctor123`, `cashier/cashier123`

### Catalog filters
- `kind=service` or `kind=goods`
- `q=` search by name or code

## Accounting
See `docs/ACCOUNTING-PATH.md`. Finance & inventory code blocked until checklist is complete.
