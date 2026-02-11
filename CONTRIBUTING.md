# Contributing to conf

Thank you for your interest in contributing!

## Development Setup

```bash
git clone https://github.com/nuln/conf.git
cd conf
make all
```

## Pull Requests

1. Fork the repo and create your branch from `main`.
2. Add tests for any new code.
3. Ensure `make all` passes.
4. Issue that pull request!

## Adding a New Codec

1. Create a new sub-package (e.g. `ini/`).
2. Implement the `conf.Codec` interface.
3. Call `conf.Register("ini", ...)` in your `init()`.
4. Import your package in `drivers/drivers.go`.
5. Add tests using `conftest.Suite`.

## Code Style

- Run `gofmt -s` and `goimports`.
- Follow [Effective Go](https://go.dev/doc/effective_go).
