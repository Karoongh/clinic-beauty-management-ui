# Accounting & audit notes

> **accounting-audit-dev**: This prototype visualizes flows. Production must pass the skill gates before writing ledger backend code.

## Prototype vs production

| Topic | Prototype | Production requirement |
|-------|-----------|------------------------|
| Opening stock | Labelled موجودی اولیه | Inventory opening journal / stock move; **not** purchase invoice |
| Sale / purchase | UI documents | Numbered documents + stock + AR/AP impact |
| Returns | Tabs exist | Linked to original doc; stock reverse; audit |
| Deposit / wallet | UI credit on accept | Liability or customer credit account; not service revenue until earned |
| Tax / discount | Footer fields | Configured rates; tax accounts; audit trail |
| Telegram confirm | Alert + local wallet | Idempotent webhook; staff user id on approval |

## Hard product rules (already decided in UI)

1. **موجودی اولیه ≠ خرید**
2. **بیعانه تأییدشده → شارژ کیف پول**; usable at settlement
3. Admission numbers start at **1** (clinic policy in UI)
4. Staff decide accept/reject **inside the app**, not inside Telegram

## Production checklist (before backend money code)

- [ ] Chart of accounts for clinic (cash, bank, AR, AP, revenue treatments, consumables, VAT, customer credits)
- [ ] Document sequences unique per clinic
- [ ] Void/credit-note rules (no silent delete)
- [ ] Stock valuation method chosen (e.g. weighted average) and documented
- [ ] Fiscal period open/close
- [ ] Who/when on every money and stock posting
- [ ] Wallet ledger sub-account per patient
- [ ] Tests: purchase→sale stock; deposit→wallet→settlement; return paths

## Visual → math map (examples)

| UI | Calculation / link |
|----|-------------------|
| Line qty × price | Line total |
| Section totals | Sum goods / sum services |
| Discount amount/% | Applied to subtotal per mode |
| Tax amount/% | Applied after discount policy |
| Wallet checkbox | `final = max(0, final - min(wallet, final))` |
| Ledger row click | Opens sale or buy document context |

## Multi-business

This product is **single clinic specialty**. Multi-tenant isolation is out of scope unless the owner expands the product later.
