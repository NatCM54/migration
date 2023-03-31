.DEFAULT_GOAL := migration_up
MAKEFLAGS += --no-print-directory
include local.env

migrate_up:
	migrate -path $(MIGRATION_FILES_PATH) -database "mysql://$(DB_MYSQL_USERNAME):$(DB_MYSQL_PASSWORD)@tcp($(DB_MYSQL_HOST):$(DB_MYSQL_PORT))/$(DB_MYSQL_DATABASE)" -verbose up $(step)

migrate_down:
	migrate -path $(MIGRATION_FILES_PATH) -database "mysql://$(DB_MYSQL_USERNAME):$(DB_MYSQL_PASSWORD)@tcp($(DB_MYSQL_HOST):$(DB_MYSQL_PORT))/$(DB_MYSQL_DATABASE)" -verbose down $(step)

migrate_fix:
	migrate -path $(MIGRATION_FILES_PATH) -database "mysql://$(DB_MYSQL_USERNAME):$(DB_MYSQL_PASSWORD)@tcp($(DB_MYSQL_HOST):$(DB_MYSQL_PORT))/$(DB_MYSQL_DATABASE)" force $(version)

migrate_create:
	migrate create -ext sql -dir $(MIGRATION_FILES_PATH) -seq $(title)

go_migrate_up:
	go run . migrate up $(step)