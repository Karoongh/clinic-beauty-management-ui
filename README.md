# Clinic Beauty Management UI

**Mobile-first Persian UI prototype** for a beauty/clinic management product (کلینیک زیبایی).

This repository is **docs-first**: read the guides under `docs/` and `AGENTS.md` before opening `prototype/`. You should not need to scan the entire JavaScript codebase to understand product rules, screens, or accounting intent.

| Item | Value |
|------|--------|
| Status | Interactive UI prototype (not production backend) |
| Language / RTL | Persian (FA), right-to-left |
| Stack | Single-page HTML + CSS + vanilla JS |
| Demo entry | [`prototype/clinic-mobile-single.html`](prototype/clinic-mobile-single.html) (all-in-one file) |
| Source modules | `prototype/index.html` · `styles.css` · `app.js` |

---

## Quick start

1. Open **`prototype/clinic-mobile-single.html`** in a mobile browser or Chrome DevTools device mode.
2. Login is prototype-only (any credentials advance past the gate).
3. Work from **خانه (Home)** and bottom navigation.

Optional local split source:

```bash
cd prototype
python3 -m http.server 8080
# open http://localhost:8080/
```

---

## What this product is

A **specialized clinic/beauty** workspace (not multi-business ERP):

- Admission (پذیرش) with channel source
- Appointments: request → confirmed reservation → admission on arrival
- Telegram (and phone / Instagram / WhatsApp) **inbound requests** handled inside the app
- Finance: sale, purchase, returns, transfer — Excel-like invoice tables
- Catalog + **item ledger** (stock movements) opening into invoices
- Opening stocktake / Excel migration (موجودی اولیه ≠ purchase)
- Patient list, profile, **wallet** (deposit credit usable at settlement)
- Analytics charts (line series, Jalali/Gregorian)
- Roles & access settings

---

## Documentation map (read in this order)

| Doc | Purpose |
|-----|---------|
| [AGENTS.md](AGENTS.md) | How AI/human agents must work on this repo |
| [docs/PRODUCT.md](docs/PRODUCT.md) | Product intent, roles, user flows |
| [docs/SCREENS.md](docs/SCREENS.md) | Screen inventory and UI behaviour |
| [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) | File layout, navigation, state patterns |
| [docs/ACCOUNTING.md](docs/ACCOUNTING.md) | Money/stock rules (accounting-audit discipline) |
| [docs/CAPABILITY-INVENTORY.md](docs/CAPABILITY-INVENTORY.md) | Keep capabilities list |
| [docs/DECISION-LOG.md](docs/DECISION-LOG.md) | Important product decisions |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Change process |

---

## Skills this project expects

1. **simple-first-dev** — mobile-first, few taps, no clutter, capability inventory before redesign
2. **safe-multi-agent-github-dev** — branch + backup before multi-file edits; Judge before large changes
3. **accounting-audit-dev** — design money/stock path + checklist **before** production ledger code

---

## Scope boundaries

**In prototype:** visual flows, demo data, localStorage samples, keyboard invoice UX.

**Not in this repo (yet):** real Telegram Bot API, server DB, double-entry journal persistence, multi-tenant auth, print PDF engine.

See `docs/ACCOUNTING.md` for what production must implement when leaving prototype stage.

---

## License

Prototype code for product development. Clarify licensing with the repository owner before commercial redistribution.
