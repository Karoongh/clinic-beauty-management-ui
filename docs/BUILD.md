# Build / bundle

## Preferred demo

`prototype/clinic-mobile-single.html` is a **self-contained** file (HTML + inlined CSS + JS) for mobile testing without a server.

## Split sources (edit these)

| File | Role |
|------|------|
| `prototype/index.html` | Markup / pages |
| `prototype/styles.css` | Styles |
| `prototype/app.js` | Behaviour |

## Regenerate single-file

From the `prototype/` directory, inline CSS and JS into one HTML file:

1. Read `index.html`, `styles.css`, `app.js`
2. Replace `<link rel="stylesheet" href="styles.css">` with a `<style>` block containing CSS
3. Replace `<script src="app.js"></script>` with a `<script>` block containing JS
4. Write `clinic-mobile-single.html` (and optionally `clinic-mobile.html`)

Open `clinic-mobile-single.html` directly in a mobile browser after bundling.
