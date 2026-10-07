# Prototype sources

## Quick path (recommended)

```bash
python3 scripts/decode-prototype.py
```

Then open **`prototype/clinic-mobile-single.html`** in a mobile browser.

Encoded sources live under `prototype/encoded/*.gz.b64` (gzip + base64) so the full UI is recoverable offline without browsing megabytes of raw HTML in git history noise.

## After decode

| File | Role |
|------|------|
| `clinic-mobile-single.html` | Full demo, no server |
| `index.html` + `styles.css` + `app.js` | Editable split sources |

See [docs/BUILD.md](../docs/BUILD.md).
