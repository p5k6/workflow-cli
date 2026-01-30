# Claude Code Guidelines

## Git Workflow
- Always use feature branches for changes, never commit directly to master/main
- Branch naming: `fix/description`, `feature/description`, `chore/description`

## Project Overview
- Go CLI client for Hephy Workflow (originally Deis Workflow) - a PaaS for Kubernetes
- Module path: `github.com/p5k6/workflow-cli`
- Forked from abandoned `teamhephy/workflow-cli` in January 2026

## Build & Test
- Requires Docker for building
- `make build` - Build dev binary for current OS (outputs to `_dist/deis`)
- `make build-stable` - Build release binaries for all platforms
- `make test` - Run tests in Docker

## Version Format (SemVer)
- Dev builds: `{TAG}-dev+g{SHA}` (e.g., `v3.0.0-dev+gbc768787`)
- Stable releases: `{TAG}` (e.g., `v3.0.0`)
- The `g` prefix on SHA indicates git commit

## Dependencies
- `github.com/p5k6/controller-sdk-go` - API client (forked from teamhephy)
- `pkg/prettyprint` - Inlined locally (was teamhephy/pkg)

## Historical Context
- Remaining `teamhephy` references in README.md and comments are intentional historical attribution
- Original project was abandoned ~2021, this fork modernized in 2026
