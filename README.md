# Go API

An API in Golang.

# Initial Setup

1. Copy env and update variables

```
cp .envrc.example .envrc
```

2. Install dependencies

```
go mod vendor
```

4. Run Migrations to init db tables

```
make db/migrations/up
```
4. Run the app

```
make run/api
```

