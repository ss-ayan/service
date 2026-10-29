# service

Self-host wrapper around `github.com/ss-ayan/ayan` (pinned by tag in `go.mod`). `cmd/service/main.go` loads `ayan.yaml` and runs `server.New`. Deploy configs are in `deploy/` (Dockerfile, compose, fly.toml, render.yaml); the Docker build needs a `gh_token` build secret because ayan is private. To upgrade: `go get github.com/ss-ayan/ayan@vX.Y.Z && go mod tidy`.
