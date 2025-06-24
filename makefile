# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/taskdb?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Команда для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Запуск приложения
run:
	go run cmd/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go


