# AGENTS.md

## Cursor Cloud specific instructions

### Overview

This is **Bubble**, a 2D breakout/brick-breaker game written in Go using the [Ebiten](https://github.com/hajimehoshi/ebiten) v1 game engine. It is a single standalone binary with no external services, databases, or APIs.

### System dependencies (pre-installed in snapshot)

Ebiten requires OpenGL/X11/ALSA dev headers: `libgl1-mesa-dev`, `xorg-dev`, `libasound2-dev`. These are installed in the VM snapshot and do not need reinstallation.

### Build & run

- **Build:** `go build -o bubble .`
- **Lint:** `go vet ./...`
- **Run:** `DISPLAY=:1 ./bubble` (requires an X display; use Xvfb if none is available)

### Gotchas

- The project originally had no `go.mod`. The update script runs `go mod init` (guarded) + `go mod tidy` to ensure dependencies are resolved.
- Ebiten v1 is used (`github.com/hajimehoshi/ebiten` without `/v2`). The resolved version is `v1.12.13`.
- The game needs a display to render. In Cloud Agent VMs, `DISPLAY=:1` is typically available. If not, start Xvfb: `Xvfb :99 -screen 0 1280x720x24 &` and use `DISPLAY=:99`.
- There are no automated tests in this repository. Validation is done via `go vet` and manual GUI testing.
