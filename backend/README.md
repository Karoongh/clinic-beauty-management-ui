# Clinic Beauty Management — Backend

Go modular monolith for the clinic/beauty product.

**Status:** Auth + Patients + Admissions ready (demo data).  
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

Demo accounts: `manager/manager123`, `reception/reception123`, `doctor/doctor123`, `cashier/cashier123`

## Patients

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/patients?q=` | Bearer | List / search |
| GET | `/api/v1/patients/:id` | Bearer | Profile |
| GET | `/api/v1/patients/:id/wallet` | Bearer | Wallet balance (read-only) |

## Admissions

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/admissions` | Bearer | All admissions |
| GET | `/api/v1/admissions/today` | Bearer | Today’s admissions |
| POST | `/api/v1/admissions` | Bearer | Create new admission |
| GET | `/api/v1/admissions/:id` | Bearer | Single admission |

### Create body example

```json
{
  "patient_id": "p1",
  "channel": "telegram",
  "service": "بوتاکس",
  "doctor": "دکتر احمدی",
  "notes": ""
}
```

Valid channels: `walk-in`, `phone`, `telegram`, `instagram`, `whatsapp`  
Admission numbers start from 1 (clinic policy).

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

1. Appointments / Requests module
2. After accounting checklist Done → finance & inventory
