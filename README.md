# test-github-actions
Testing github actions


## Test goreleaser locally

```shell
go install github.com/goreleaser/goreleaser@latest
goreleaser release --snapshot --clean --release-notes CHANGELOG.md
```

## Create new release

```shell
go changie batch minor
go changie merge
```