# Clinic Beauty Management — Backend

Go modular monolith for the clinic/beauty product.

**Status:** Auth + Patients + Admissions + Appointments ready (demo data).  
No production money/stock mutation yet.

## Quick start

```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

## Auth
`POST /api/v1/auth/login` · `GET /api/v1/auth/me`  
Demo: manager/manager123 · reception/reception123 · doctor/doctor123 · cashier/cashier123

## Patients
`GET /api/v1/patients?q=` · `GET /api/v1/patients/:id` · `GET /api/v1/patients/:id/wallet`

## Admissions
`GET /api/v1/admissions` · `GET /api/v1/admissions/today` · `POST /api/v1/admissions` · `GET /api/v1/admissions/:id`

## Appointments / Requests

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/appointments/requests` | List with filters `?channel=&status=&type=` |
| GET | `/api/v1/appointments/requests/:id` | Single request |
| POST | `/api/v1/appointments/requests/:id/decide` | Accept or reject |

### Filters
- channel: `telegram` · `phone` · `instagram` · `whatsapp`
- status: `pending` · `accepted` · `rejected`
- type: `reserve` · `deposit`

### Decide body
```json
{ "action": "accept" }
```
or `{ "action": "reject", "notes": "..." }`

Note: accepting a deposit currently only changes status. Wallet credit will be added under the accounting gate.

## Accounting rules
See `docs/ACCOUNTING-PATH.md`. No finance code until checklist is complete.
