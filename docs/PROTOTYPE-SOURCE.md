# Full interactive prototype sources

## Docs-first repo

This repository prioritizes **product documentation** so agents and developers do not need to read all UI code first.

## Running the UI

### Option A — decode from git (when encoded chunks are complete)

```bash
git clone https://github.com/Karoongh/clinic-beauty-management-ui.git
cd clinic-beauty-management-ui
python3 scripts/decode-prototype.py
# open prototype/clinic-mobile-single.html
```

### Option B — project artifacts (authoring workspace)

During active design, the live single-file demo is also maintained as:

- `clinic-mobile-single.html` in the product workspace artifacts

Copy that file into `prototype/` if you are syncing from the design session.

### Option C — split sources after decode

Edit `index.html`, `styles.css`, `app.js`, then rebuild the single file (see [BUILD.md](BUILD.md)).

## Skills required for further work

- simple-first-dev
- safe-multi-agent-github-dev
- accounting-audit-dev (before money backend)
