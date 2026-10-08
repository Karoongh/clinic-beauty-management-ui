# Finance module

**Status:** Structure only. No production code for invoices, wallet, or payments yet.

## Accounting-audit-dev gate (must pass before any code)

Before writing handlers or repositories for this module:

1. Complete path design (documents, journals, stock impact, void rules)
2. Critique & fix
3. Full checklist (document sequences, soft-void, audit columns, wallet as liability)
4. Visual-item map from UI to calculations
5. Accounting correctness + audit-trail gates

Hard product rules (already decided):

- Opening stock is **not** a purchase
- Confirmed deposit credits **patient wallet** (liability), not revenue
- Soft-void only; never hard-delete money documents
- Document numbers unique per clinic
- Stock valuation: Weighted Average

See root `docs/ACCOUNTING.md`.
