### Install sqlc

### MacOs

`brew install sqlc`

### Ubuntu

`sudo snap install sqlc`

### Go Install

#### Go >= 1.17:

`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

#### Go < 1.17:

`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

### Install Docker Desktop

[Docker Desktop Installation Instructions](https://docs.docker.com/desktop/)

### install golang-migrate

```$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz

```

### install mockgen

```
placeholder for commmand
```

### How do I develop with this?

- If you add/remove any "queries" within the query/ directory, you should run the `make queryupdate` command, which will add update any sqlc generated code and create the necessary mocked functions for additional/existing testing.
