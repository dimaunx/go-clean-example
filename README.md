# Simple goland CLEAN architecture implementation

## Prerequisites

- go1.17.x recommended installing with [gvm]
- [curl]
- [golangci-lint] or make install/lint
- [goimports]
- [docker]
- [docker-compose]

---
**NOTE!**

On MAC M1/Intel chip requires brew and GNU MAKE upgrade

### Brew install if not installed already

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### GNU make upgrade

```bash
brew install make
export PATH="/usr/local/opt/make/libexec/gnubin:$PATH" # Add to .bashrc or .zshrc
make -version # Should be higher then 3.8.1
```

---

### List targets and help

```bash
make help
```

### Run vendor

```bash
make vendor
```

### Run linter locally

```bash
make lint
```

### Run server locally

```bash
make run
```

---
Endpoints:

server - http://localhost:8000

Auth: Header: X-Api-Key test

redis-commander - http://localhost:8903

redis - localhost:8902

<!--links-->

[goimports]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports

[docker]: https://docs.docker.com/get-docker/

[curl]: https://curl.se/download.html

[golangci-lint]: https://golangci-lint.run/usage/install/

[gvm]: https://github.com/moovweb/gvm

[docker-compose]: https://docs.docker.com/compose/install/