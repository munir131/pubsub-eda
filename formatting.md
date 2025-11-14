# Formatting & Style Guide

Last updated: 2025-11-14 12:41 (IST)

This repo may host multiple languages; keep style consistent and automated.

## General Principles
- Run formatters before every commit.
- Keep functions small and pure where practical.
- Prefer explicitness over cleverness; optimize for readability.

## Markdown
- Use sentence case for headings.
- Wrap lines at ~100 chars where convenient.
- Prefer fenced code blocks with language hints.

## Git Hygiene
- Conventional Commits style:
  - feat: add producer retry logic
  - fix: handle nil payload in consumer B
  - docs: update architecture diagram
  - chore: bump dependency versions
  - test: add integration tests for DLQ
- Keep commits focused; avoid mixing refactors with features.

## Python (if used)
- Formatter: black (line length 100)
- Linter: ruff or flake8
- Typing: mypy on changed files if practical
- Docstrings: Google style or reST

## Node/TypeScript (if used)
- Formatter: prettier
- Linter: eslint (airbnb or recommended)
- TS: strict mode, path aliases via tsconfig

## Go (if used)
- Formatter: gofmt/goimports
- Linter: golangci-lint with default linters
- Module: go 1.22+ suggested

## Shell
- Use bash with `set -euo pipefail`
- Validate inputs; prefer long flags

## YAML/JSON
- 2 spaces indentation
- Keep keys sorted where it aids diff readability

## PR Checklist
- [ ] Code formatted and linted
- [ ] Tests added/updated and passing
- [ ] Docs updated (if behavior changes)
- [ ] Breaking changes called out
