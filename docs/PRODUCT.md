# Product guide — Clinic Beauty Management

## Intent

Specialized **beauty/clinic** operations software (فارسی، موبایل‌اول).  
Not a generic multi-business ERP. Focus: admissions, appointments, treatments catalog, stock of consumables, invoices, patient wallet, inbound messaging channels.

## Primary roles

| Role | Typical access |
|------|----------------|
| Manager | All modules + users/settings |
| Reception | Admission, patients, appointments, catalog |
| Doctor | Patients, appointments, catalog (limited finance) |
| Cashier | Sale/purchase, reports |
| Warehouse | Catalog, stocktake, purchase |

(Prototype: Settings → Users & access.)

## Core flows

### 1. Appointment lifecycle

```
Request (Telegram / phone / Instagram / WhatsApp / walk-in)
    → Pending decision in app
    → Accept → Reservation (confirmed time; no admission code yet)
    → Patient arrives → Admission (acceptance number from 1)
```

Staff **do not need to open Telegram**. The bot is a channel; decisions stay in the app; replies go out via bot.

### 2. Admission

- Number starts at 1 (prototype).
- **Channel** dropdown: walk-in, phone, telegram, instagram, whatsapp.
- If the path originated from a **Telegram** request, channel is auto-selected as Telegram.

### 3. Finance documents

Tabs: Sale · Purchase · Sale return · Purchase return · Transfer (حواله).

Excel-like body: row / name / unit / qty / price / line total / note / delete.
Goods and services sections separated.
Footer: subtotal, discount, discount code, tax, wallet use, final amount — left-aligned stack.

### 4. Deposit → wallet

1. Inbound request type **بیعانه (deposit)**
2. Staff confirms receipt
3. Amount credits **patient wallet**
4. Balance visible on patient profile / table
5. At sale settlement, optional “use wallet”
6. Product intent: balance also shown in the **customer’s Telegram bot** (production)

Deposit is **credit**, not recognition of service revenue until treatment is invoiced per accounting policy.

### 5. Opening stock (انبارگردانی)

On first use (or Settings → Stocktake):

- Import patients Excel
- Import goods + quantities Excel
- Or manual opening lines

**Opening stock is not a purchase.** It appears in item ledger as موجودی اولیه.

### 6. Catalog ledger

Services & goods → item summary → full movement ledger (date range) → click movement → **Finance** opens with that invoice context.

## Non-goals (current product boundary)

- Multi-company holding / multi-trade shops in one tenant
- Full double-entry UI in the prototype (see ACCOUNTING.md for production target)
- Payroll, complex manufacturing
