# go-version
Set version, date and git commit in binary

Example

```
go build -ldflags="-X 'github.com/sparmcgit/go-version.Version=1.0.0' \
        -X 'github.com/sparmcgit/go-version.BuildTime=$(date)' \
        -X 'github.com/sparmcgit/go-version.CommitHash=$(git rev-parse HEAD)'" \
        -o main_version .
```

