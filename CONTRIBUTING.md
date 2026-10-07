# Contributing

## Before you start

1. Read [AGENTS.md](AGENTS.md).
2. For UI: follow **simple-first-dev** (mobile-first, no feature drop without Keep/Drop).
3. For money/stock: follow **accounting-audit-dev** gates before backend work.
4. For GitHub multi-file work: follow **safe-multi-agent-github-dev** (branch + backup).

## Workflow

1. Branch from `main`: `feature/...` or `fix/...`
2. Edit `prototype/index.html`, `styles.css`, `app.js` as needed
3. Rebuild single-file demo if you distribute it
4. Update the relevant file under `docs/` in the **same PR**
5. Open a Pull Request; do not force-push `main`

## PR checklist

- [ ] Docs updated if behaviour changed
- [ ] Mobile checked
- [ ] No secrets committed
- [ ] Accounting rules not contradicted (see docs/ACCOUNTING.md)
