# AGENTS.md — Clinic Beauty Management UI

Instructions for **human developers and AI agents** working on this repository.

## 0. Read before code

1. [docs/PRODUCT.md](docs/PRODUCT.md) — what the product is
2. [docs/SCREENS.md](docs/SCREENS.md) — where each feature lives
3. [docs/ACCOUNTING.md](docs/ACCOUNTING.md) — money/stock rules
4. [docs/DECISION-LOG.md](docs/DECISION-LOG.md) — frozen decisions

Do **not** invent multi-business ERP behaviour. This product is **clinic/beauty only**.

## 1. Skills (mandatory)

| Skill | When |
|-------|------|
| **simple-first-dev** | Any UI change, new screen, navigation, mobile layout |
| **safe-multi-agent-github-dev** | Multi-file or risky GitHub edits: branch, backup, critique, Judge |
| **accounting-audit-dev** | Before any production code for invoices, stock, payments, wallet, tax, reports |

### simple-first-dev rules (summary)

- Mobile-first; one primary job per view
- Capability Inventory on existing code before dropping features
- Surface every **Keep** capability in the next prototype
- Prefer cards/sheets over deep menus

### safe-multi-agent-github-dev rules (summary)

- New branch from default; no force-push without explicit user confirmation
- Backup branch/tag before large changes
- Prefer PR over direct main commits when possible

### accounting-audit-dev rules (summary)

- Design path + checklist **before** backend money code
- Opening stock ≠ purchase
- Deposit → wallet credit, not final revenue for the service
- Document number sequences, void/soft-delete, audit who/when in production

## 2. Repository layout

```
prototype/
  clinic-mobile-single.html   # preferred demo (bundled)
  index.html                  # structure
  styles.css                  # styles
  app.js                      # behaviour + demo data
docs/                         # product & agent guides
AGENTS.md
README.md
CONTRIBUTING.md
```

Prefer editing **split sources** (`index.html`, `styles.css`, `app.js`) then regenerating the single-file bundle for distribution.

## 3. Change policy

| Change type | Required |
|-------------|----------|
| Visual / copy only | Update `docs/SCREENS.md` if behaviour changes |
| New screen / flow | Update PRODUCT + SCREENS + DECISION-LOG |
| Money / stock / tax / wallet | Pass accounting-audit gates; update ACCOUNTING.md |
| Drop a feature | Explicit user Keep/Drop — never silent drop |

## 4. Demo data & localStorage

Prototype keys (non-exhaustive):

- `clinic-users-access` — roles/permissions
- `clinic-wallets` — patient wallet balances
- `clinic-stocktake-done` / `clinic-opening-stock`
- `clinic-stocktake-skip`
- `adm-from-tg` (session) — admission channel from Telegram

Do not store real secrets (Bot tokens are UI placeholders only).

## 5. Coding conventions

- Persian UI strings; RTL (`dir="rtl"`)
- Prefer event delegation for dynamic lists
- Invoice lines: goods/services sections, Enter-driven pickers
- Avoid new global frameworks unless the owner requests them

## 6. Testing checklist (prototype)

- [ ] Mobile width ≤ 390px usable
- [ ] Login → Home
- [ ] Requests filters + accept/reject
- [ ] Deposit accept credits wallet
- [ ] Ledger row → Finance tab with line filled
- [ ] Stocktake page reachable
- [ ] Patient row / recent list opens profile
- [ ] Dark theme if present still readable

## 7. Out of scope for casual PRs

- Real Telegram webhook server
- PostgreSQL/SQLite production schema (design in ACCOUNTING.md first)
- Changing product from clinic-only to multi-trade without owner approval
