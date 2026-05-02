DB_URL=postgres://postgres:w5bf3xxx@localhost:5432/marketplace?sslmode=disable

migrate-up:
	migrate -path ./db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./db/migrations -database "$(DB_URL)" down $(step)

migrate-force:
	migrate -path ./db/migrations -database "$(DB_URL)" force $(version)

migrate-version:
	migrate -path ./db/migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir ./db/migrations -seq $(name)