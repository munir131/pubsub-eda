# Restricted Files & Guardrails

Last updated: 2025-11-14 12:41 (IST)

This list guides AI/codegen tools on what NOT to modify without explicit approval.

## Do not commit secrets
- .env, .env.local, any files containing API keys/passwords
- Cloud credentials, kubeconfigs, SSH keys

## Avoid automated edits to
- Lockfiles: `package-lock.json`, `yarn.lock`, `poetry.lock`, `Pipfile.lock`, `go.sum`
  - Rationale: Keep dependency changes explicit and reviewed.
- Binary assets: images, PDFs, diagrams exports
- Generated code: any files under `gen/` or marked "GENERATED - DO NOT EDIT"
- Large data files: anything in `data/` over 5MB

## Protected config
- CI/CD workflows: `.github/workflows/*`, unless task explicitly requests it
- Infrastructure templates: `docker-compose.yml`, `Dockerfile*`, `infra/*` (review changes carefully)

## Documentation invariants
- Keep `architecture.md` source-of-truth; diagrams should reference it.
- Do not rewrite changelogs/history sections; only append.

## Safe areas for AI edits
- Source code under `producers/`, `consumers/`, small helper scripts in `scripts/`
- Tests under `tests/`
- Docs under `docs/` and top-level *.md files (except restricted sections above)

## Process
- If a change touches a restricted file, open a PR with clear rationale.
- Include before/after summary and validation steps.
