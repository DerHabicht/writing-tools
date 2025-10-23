# writing-tools

## Project Standards

- [Semantic Versioning](https://semver.org/spec/v2.0.0.html)
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)

## Testing

Some of the tests will attempt to invoke XeLaTeX. To skip these, invoke with the `-s` flag:

```
go test -s ./...
```