# Decision log

| Date | Decision | Rationale |
|------|----------|-----------|
| 2026-09/10 | Clinic-only product (not multi-trade ERP) | Owner focus on treatment domain |
| 2026-10 | Request → Reservation → Admission on arrival | Code only when patient enters |
| 2026-10 | Staff never required to open Telegram | Bot is channel; app is control plane |
| 2026-10 | Opening stock ≠ purchase | Migration / stocktake integrity |
| 2026-10 | Deposit confirms credit wallet | Usable at settlement; visible in bot (prod) |
| 2026-10 | Request filters: همه first; default status همه | Clearer list UX |
| 2026-10 | Pending cards pulse + icon ring | Visibility of new work |
| 2026-10 | Accepted not dimmed; rejected gray | Status readability |
| 2026-10 | Ledger click → full finance page | Not only a mini popup |
| 2026-10 | Invoice foot left stack + stronger table contrast | Reduce visual confusion |
| 2026-10 | Docs-first separate GitHub repo | Agents avoid reading all code |

Append new decisions; do not silently overwrite history.
