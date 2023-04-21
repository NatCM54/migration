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

```golang
func(cmd *cobra.Command, args []string) {
  action := args[0]
  step, _ := strconv.Atoi(args[1])
  m := lib.Migrate()

  if action == "up" {
   m.Steps(step)
   fmt.Println("Running migrate up command")
  } else if action == "down" {
   m.Steps(step * (-1))
   fmt.Println("Running migrate down command")
  }
 },
```
