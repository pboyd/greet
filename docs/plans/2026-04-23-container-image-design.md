# Container Image & GHCR Publish Design

## Overview

Build a minimal container image for the `hello` HTTP server and publish it to GHCR via GitHub Actions.

## Dockerfile

Multi-stage build:

- **Stage 1** — `golang:1.26-bookworm`: compile a fully static binary (`CGO_ENABLED=0`, `GOOS=linux`)
- **Stage 2** — `scratch`: copy only the binary; no shell, no CA certs, no additional layers

The resulting image contains exactly one file.

## GitHub Actions Workflow

Single workflow file at `.github/workflows/docker.yml`.

**Triggers:**

| Event | Result |
|---|---|
| Push to `main` | `ghcr.io/pboyd/hello:latest` |
| Push tag `v*` | `ghcr.io/pboyd/hello:1.2.3`, `1.2`, `1` |

`latest` tracks `main` only — it is not updated on tagged releases.

**Steps:**

1. Checkout
2. Set up Docker Buildx
3. Log in to GHCR using `GITHUB_TOKEN` (no additional secrets required)
4. Derive tags and labels via `docker/metadata-action`
5. Build and push via `docker/build-push-action`

## Decisions

- **scratch over distroless** — app uses only stdlib, makes no outbound HTTPS calls, no need for CA certs or timezone data
- **Single workflow** — `docker/metadata-action` handles the tagging matrix cleanly; no duplication
- **latest = main, not latest release** — explicit user requirement
