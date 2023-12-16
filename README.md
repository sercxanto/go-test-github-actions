# test-github-actions
Testing github actions


## Developer documentation

### Start with a new change

Call `changie new`:

```shell
changie new
```

This will ask for the kind of change and create a new file in `./changes/unreleased`.

## Create new release

The release process consists of the following steps:

1. Create changelog locally
2. Test goreleaser locally
3. Tag the release locally and trigger goreleaser on Github CI

### Create changelog locally

`changie batch` collects unreleased changes info from `./changes/unreleased` and creates a
new version file like `./changes/v1.2.3.md`.

`changie merge` collects version files from the `./changes` folder and updates `CHANGELOG.md`.

Change `minor`to the type of change:

```shell
changie batch minor
changie merge
```
You may want to call the changie commands with the `--dry-run` to preview the changelog.

Don't forget to commit the changes so that the workspace is clean:

```shell
git add .
git commit ...
```

### Test goreleaser locally

```shell
go install github.com/goreleaser/goreleaser@latest
goreleaser release --snapshot --clean --release-notes .changes/$(changie latest).md
```

### Tag the release locally and trigger goreleaser on Github CI

```shell
git tag $(changie latest)
git push origin main && git push --tags
```
