# **How to run?**

First, start setup local.env

``` .env
DB_MYSQL_USERNAME=abserp-dev
DB_MYSQL_PASSWORD=aa11111
DB_MYSQL_HOST=localhost
DB_MYSQL_PORT=25060
DB_MYSQL_DATABASE=abserp-dev
MIGRATION_FILES_PATH=database/migrations
```

### Migrate

1. migrate step 1 UP

```bash
go run . migrate up ...("step")
```

2. migrate step 1 DOWN

```bash
go run . migrate down ...("step")
```

