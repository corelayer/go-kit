# Run example
```shell
go run -ldflags="-X main.VersionCommit=$(git rev-list --abbrev-commit --abbrev=8 -1 HEAD) -X main.VersionSemVer=$(git name-rev HEAD --name-only)" examples/app/main.go version
```

# Build example

```shell
go build -o build/bin/app -ldflags="-X main.VersionCommit=$(git rev-list --abbrev-commit --abbrev=8 -1 HEAD) -X main.VersionSemVer=$(git name-rev HEAD --name-only)" examples/app/main.go
```