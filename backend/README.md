# Clinic Beauty Management — Backend

Go modular monolith. **Auth + Patients use PostgreSQL.**  
No production money/stock mutation.

## Quick start

```bash
cd backend
docker compose -f deploy/compose/docker-compose.yml up --build
```

On first start, Postgres runs migrations from `migrations/` automatically.

If you change migrations after the volume already exists:

```bash
docker compose -f deploy/compose/docker-compose.yml down -v
docker compose -f deploy/compose/docker-compose.yml up --build
```

## Demo logins (seeded)

All seeded users currently share password: **`password`**

| Username | Role |
|----------|------|
| manager | manager |
| reception | reception |
| doctor | doctor |
| cashier | cashier |

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"manager","password":"password"}'
```

## Modules

| Module | Storage | Notes |
|--------|---------|-------|
| Auth | PostgreSQL | bcrypt |
| Patients | PostgreSQL | list / profile / wallet read |
| Admissions | in-memory demo | next to migrate |
| Appointments | in-memory demo | next to migrate |
| Catalog | in-memory demo | next to migrate |
| Analytics | static demo | next to real queries |

## Accounting gate

Finance and inventory **must not** be implemented until `docs/ACCOUNTING-PATH.md` checklist is fully Done.

## Safety

- Branch: `feat/db-users-patients`
- Safety: `safety/20261008-db-users-patients`
