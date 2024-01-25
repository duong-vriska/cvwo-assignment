DB_PATH = postgresql://root:nmai1202@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local

migration_up: migrate -path db/migrations -database "${DB_PATH}" -verbose up

migration_down: migrate -path db/migrations -database "${DB_PATH}" -verbose down

migration_fix: migrate -path db/migrations -database "${DB_PATH}" force VERSION
