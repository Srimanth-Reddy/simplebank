# Simple Bank

A Go backend project built with **Gin** & **gRPC**, using **Postgres** and **Redis**.  
Designed to run in **Docker/Kubernetes** on **AWS** with a full **CI/CD pipeline**.

---

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Go 1.21+](https://go.dev/doc/install)
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

On macOS, you can install the tools via Homebrew:

```
brew install sqlc
brew install golang-migrate
```

Run postgres, createdb, and migrateup commands in the makefile:
```
make postgres
make createdb
make migrateup
```

Once done with the setup, you can play around with the code. Cheers!