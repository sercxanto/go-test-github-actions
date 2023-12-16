# test-github-actions
Testing github actions


## Test goreleaser locally

Exchange `v1.2.3.md` with the upcoming version

```shell
go install github.com/goreleaser/goreleaser@latest
goreleaser release --snapshot --clean --release-notes .changes/$(changie latest).md
```

## Start with a new change

Call `changie new`:

```shell
changie new
```

This will ask for the kind of change and create a new file in `./changes/unreleased`.

## Create new release

Change `minor`to the type of change:

```shell
changie batch minor
changie merge
```
