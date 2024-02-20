BINARY_NAME=simpbb
GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING=file:./simpbb.db
GOOSE_MIGRATION_DIR=./internal/migrations

build:
	go build -o ./bin/${BINARY_NAME} ./cmd/

air:
	air -c .air.toml
	npm install
	go generate ./assets/tailwind.go

start: build
	./bin/${BINARY_NAME}

clean: 
	go clean
	rm -rf node_modules
	rm -f ./bin/${BINARY_NAME}

migrate-up:
	go run ./cmd migrate up

migrate-down:
	go run ./cmd migrate down