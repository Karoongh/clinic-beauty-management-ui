# Architecture (prototype)

## Runtime

- No build step required for `clinic-mobile-single.html`.
- Split mode: `index.html` + `styles.css` + `app.js`.
- Navigation: `go(pageId)` toggles `.page.active`, bottom nav, side nav.

## Important JS areas (app.js)

| Area | Responsibility |
|------|----------------|
| Navigation / sheets | `go`, more menu, scroll helpers |
| Chart | Search-Console-style multi series |
| Patients | `PATIENTS_DATA`, `renderPatients`, profile, wallet |
| Finance lines | `lineState`, pickers, `renderLines`, `updateSums` |
| Stocktake / import | CSV templates, opening stock flags |
| Catalog / ledger | `CATALOG_UI`, `LEDGER_DEMO`, `openLedgerInvoice` |
| Telegram requests | `TG_REQUESTS`, filters, accept/reject |
| Access control UI | `USERS`, `PERMS`, role presets |

## State

- In-memory demo arrays
- `localStorage` for wallets, users, stocktake flags
- `sessionStorage` for Telegram → admission channel handoff

## UI patterns

- Cards + bottom sheets for pickers/summaries
- Excel-like tables for invoices
- Chip filters for request lists
- Dynamic DOM for item/value pickers (avoid sticky orphan sheets)

## Production evolution (suggested)

```
Browser UI  →  API (auth, clinic tenant)  →  DB
                ↘ Telegram webhook worker
                ↘ Print / Excel export service
```

Do not hard-code production secrets in the frontend.
