# bun-driver

## Install

```
$ go get github.com/innotechdevops/bun-driver
```

## How to use

- MariaDB

```go
driver := bundriver.NewMariaDBDriver(&bundriver.MariaConfig{
    User:         os.Getenv("MARIA_USER"),
    Pass:         os.Getenv("MARIA_PASS"),
    Host:         os.Getenv("MARIA_HOST"),
    DatabaseName: os.Getenv("MARIA_DATABASE"),
    Port:         bundriver.MariaDefaultPort,
    Loc:          "Asia%2FBangkok",
	PoolOptions:  &dbre.DbPoolOptions{
        MaxIdleConns:    2,
        MaxOpenConns:    10,
        ConnMaxLifetime: 30 * time.Minute,
    },
})
```

- PostgreSQL

```go
driver := bundriver.NewMariaDBDriver(&bundriver.PostgresConfig{
    SSLMode: "disable",
    Config: &bundriver.Config{
        User:         os.Getenv("POSTGRES_USER"),
        Pass:         os.Getenv("POSTGRES_PASS"),
        Host:         os.Getenv("POSTGRES_HOST"),
        DatabaseName: os.Getenv("POSTGRES_DATABASE"),
        Port:         bundriver.PostgresDefaultPort,
        Loc:          "Asia%2FBangkok",
        PoolOptions:  &dbre.DbPoolOptions{
            MaxIdleConns:    2,
            MaxOpenConns:    10,
            ConnMaxLifetime: 30 * time.Minute,
        },
    }
})
```