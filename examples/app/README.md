# Run example

```shell
go run -ldflags="-X main.VersionCommit=$(git rev-list -1 HEAD) -X main.VersionSemVer=$(git name-rev HEAD --name-only)" examples/app/main.go version
```