# test-github-actions
Testing github actions


## Test goreleaser locally

```shell
go install github.com/goreleaser/goreleaser@latest
goreleaser release --snapshot --clean --release-notes CHANGELOG.md
```

## Start with a new change

Call `changie new`:

```shell
changie new
```

This will ask for the kind of change and create a new file in `./changes/unreleased`.

## Create new release

```shell
changie batch minor
changie merge
```

