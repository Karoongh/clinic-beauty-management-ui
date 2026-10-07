# Screen inventory

## Shell

| Element | Behaviour |
|---------|-----------|
| Login | Prototype gate |
| Top title | Current page name |
| Bottom nav | Home, Admission, More |
| More sheet | Patients, Appointments, Finance, Services, Stocktake, Settings, … |
| Theme | Light/dark if toggled in UI |

## Home (`p-home`)

- KPI cards (tap → scroll to related section / chart)
- Chart (line series; category: goods / services / revenue / admissions)
- Telegram requests preview (pending)
- Recent admissions

## Admission (`p-admission`)

- New admission form: number, date, patient, **channel**, service, doctor, notes
- Today’s list
- Print placeholder

## Patients (`p-patients`)

- Wallet explanation card
- **Recent patients** — click opens **profile**
- Full table: #, code, name, age, gender, mobile, visits, invoices, total, **wallet**, account status
- Search, sort, Excel export placeholder
- Profile: fields + shortcuts to admission / sale

## Appointments (`p-appointments`)

Tabs: Requests · Reservations · Today

### Requests filters

- Channel: all, telegram, phone, instagram, whatsapp
- Status: all, pending, accepted, rejected (default **all**; **همه** first in every chip row)
- Type: all, reserve, deposit
- Date from–to
- Reset filters

### Request card visuals

- Pending: stronger border pulse + channel icon “ring”
- Accepted: green, **not** faded
- Rejected: gray border/text
- Channel icons: Telegram, Instagram, WhatsApp, Phone

## Finance (`p-finance`)

Tabs: sale, buy, sale return, purchase return, transfer.

Invoice structure:

1. Head (meta)
2. Body (table, horizontal scroll if needed)
3. Foot (totals) — clearer contrast, foot stack toward **visual left**

Keyboard flow (sale): item → qty → price pickers via Enter (see app.js).

## Services & goods (`p-services`)

- Search per tab
- Click item → summary sheet (stock, last buy/sale)
- More → ledger with date range → click row → finance invoice

## Stocktake (`p-stocktake`)

Migration + manual opening stock.

## Settings (`p-settings`)

Calendar Jalali/Gregorian, clinic info, Telegram bot fields (placeholder), **users & permissions**, link to stocktake, logout.
