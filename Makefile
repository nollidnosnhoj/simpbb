BINARY_NAME=simpbb
MAIN_CMD_FILEPATH=./cmd

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ${MAIN_CMD_FILEPATH}
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ${MAIN_CMD_FILEPATH}

air:
	air -c .air.toml
	npm install
	go generate ./assets/tailwind.go

start: build
	./bin/${BINARY_NAME}-darwin start

start-linux: build
	./bin/${BINARY_NAME}-linux start

clean: 
	go clean
	rm -rf node_modules
	rm -f ./bin/${BINARY_NAME}-darwin
	rm -f ./bin/${BINARY_NAME}-linux