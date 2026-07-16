# Holonet

Tick-based MMO game server written in Go. A Godot client communicates with this
server over WebSockets. The server is authoritative: game state advances in a
fixed tick loop (goroutine). Design is early and evolving — currently a single
`main` package.

## Stack

- **Server:** Go, WebSocket API for the Godot client, tick-based game loop
- **Persistence:** Postgres (planned); Redis for caching (planned)
- **Blobs:** Cloudflare R2 if needed
- **Packaging/deploy:** Docker; local dev → self-hosted Dokploy PaaS (test) → production
- **Secrets:** self-hosted Infisical in dev; `.env` is local-only and gitignored

## Commands

- `just check` — format, lint (golangci-lint), and test. Run before finishing any task.
- `just run` / `just test` — dev server / tests only.
- Global tools come from mise (`mise.toml`); project dependencies via `go get` (`go.mod`).

## Conventions

- **Config:** env vars parsed once at startup in `env.go` (`parseEnv`). Defaults:
  `ENVIRONMENT=dev`, `PORT=8080`. Valid environments: `dev`, `test`, `prod`.
- **Errors:** return wrapped errors (`fmt.Errorf` with `%w`). Log-and-exit only in
  `main`; inner functions never call `os.Exit` or log errors they also return.
- **Logging:** `slog` with key-value attributes (`slog.Info("msg", "key", val)`),
  never values interpolated into the message string. Tint handler in dev.
- **Tests:** stdlib `testing` only, `t.Setenv` for env manipulation. No frameworks.
- **Formatting:** `golangci-lint fmt` (gofmt-compatible). Never hand-format.
- Prefer simple, stable, idiomatic Go and the standard library over clever
  abstractions or new dependencies. Justify any new dependency.
- The maintainer is learning Go — when using a non-obvious idiom, briefly explain
  it in your summary.

## Commit conventions

Single-line commits, `type: description` — lowercase, no scope, no period.
Types: `feat` `fix` `docs` `style` `perf` `test` `build` `ci` `chore` `revert`.

VCS is jj (colocated with git).

## Don'ts

- Never commit, branch, or otherwise mutate jj/git state — the user handles VCS.
- Never edit `.env` (real secrets). `example.env` is the committed template; keep
  it in sync when adding config.
- Keep `CHANGELOG.md` updated under `[Unreleased]` — highlight/summary level only,
  not a list of every change.
