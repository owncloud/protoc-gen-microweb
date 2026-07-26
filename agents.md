# agents.md -- protoc-gen-microweb

## Repository Overview

Protocol Buffers compiler plugin that generates HTTP REST handlers from protobuf service definitions. Written in Go. Licensed under Apache-2.0.

## Architecture & Key Paths

- `main.go` -- Plugin entry point
- `microweb.go` -- Core code generation logic
- `examples/` -- Example protobuf services and generated output
- `go.mod` / `go.sum` -- Go module definition
- `tools.go` -- Tool dependencies
- `test.sh` -- Test script

## Development Conventions

- Go codebase
- Protoc plugin protocol
- Chi router for generated handlers

## Build & Test Commands

```bash
go install                    # Install the plugin
bash test.sh                  # Run tests
```

## Important Constraints

- Licensed under Apache-2.0 (already at the OSPO target license). The broader ownCloud organization is migrating other repositories from copyleft licenses to Apache 2.0.
- Used in the oCIS build pipeline for REST handler generation.
- All contributions require a DCO sign-off.


## OSPO Policy Constraints

### GitHub Actions
- **Only** use actions owned by `owncloud`, created by GitHub (`actions/*`), verified on the GitHub Marketplace, or verified by the ownCloud Maintainers.
- Pin all actions to their full commit SHA (not tags): `uses: actions/checkout@<SHA> # vX.Y.Z`
- Never introduce actions from unverified third parties.

### Dependency Management
- Dependabot is configured for automated dependency updates.
- Review and merge Dependabot PRs as part of regular maintenance.
- Do not introduce new dependencies without discussion in an issue first.

### Git Workflow
- **Rebase policy**: Always rebase; never create merge commits. Use `git pull --rebase` and `git rebase` before pushing.
- **Signed commits**: All commits **must** be PGP/GPG signed (`git commit -S -s`).
- **DCO sign-off**: Every commit needs a `Signed-off-by` line (`git commit -s`).
- **Conventional Commits & Squash Merge**: Use the [Conventional Commits](https://www.conventionalcommits.org/) format where the repository enforces it. Many repos use squash merge, where the PR title becomes the commit message on the default branch — apply Conventional Commits format to PR titles as well. A reusable GitHub Actions workflow enforces this.

## Context for AI Agents

This is a protoc plugin. It reads protobuf service definitions with gRPC HTTP annotations and generates Go code that wires up Chi router handlers with JSON serialization. The generated code bridges protobuf services to HTTP REST endpoints.
