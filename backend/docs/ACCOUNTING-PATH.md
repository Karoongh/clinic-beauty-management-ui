# Accounting Path Design (must pass before any finance code)

This document satisfies the accounting-audit-dev gates for the clinic beauty product.

## Hard rules (already decided in UI)

1. Opening stock ≠ purchase
2. Confirmed deposit → patient wallet credit (liability), not service revenue
3. Soft-void only for money and stock documents
4. Document numbers unique per clinic
5. Stock valuation method: Weighted Average

## Document → Stock / Money impact

| Document | Stock effect | Money / account effect | Void rule |
|----------|--------------|------------------------|-----------|
| Opening Stock | +qty (type=opening) | none | correction document only |
| Purchase | +qty | AP increase | soft-void + reverse stock & AP |
| Sale | -qty | AR or cash + revenue | soft-void + reverse stock & revenue |
| Sale Return | +qty | reverse AR/revenue | must link to original sale |
| Purchase Return | -qty | reverse AP | must link to original purchase |
| Transfer | move between locations | none | soft-void |
| Deposit | none | +wallet (liability) | correction only |
| Wallet use on sale | none | -wallet, reduce payable | follows sale void |

## Checklist (gate)

### Data model
- [ ] documents + document_lines
- [ ] stock_movements (with type, qty, unit_cost, running avg)
- [ ] wallet_ledger (per patient)
- [ ] unique (clinic_id, doc_type, number)
- [ ] audit columns: created_by, created_at, voided_by, voided_at, void_reason
- [ ] soft-void flag (is_voided)

### Lifecycle
- [ ] create path
- [ ] void path with full reversal
- [ ] return path linked to original
- [ ] partial return support
- [ ] wallet debit only up to available balance

### Stock & valuation
- [ ] Weighted Average calculation documented and single source of truth
- [ ] negative stock policy defined (block or allow)
- [ ] unit of measure consistency

### Money & tax
- [ ] discount then tax then wallet order
- [ ] rounding rules documented
- [ ] AR/AP impact clear

### Tests required before merge of finance code
- [ ] purchase → sale stock correct
- [ ] deposit → wallet credit → use on sale
- [ ] void sale restores stock and wallet
- [ ] duplicate document number rejected under concurrency

## Status

Path designed. Checklist not yet marked Done.
No production finance/inventory code may be written until every applicable item is Done and reviewed.
