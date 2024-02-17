BINARY_NAME=simpbb
MAIN_CMD_FILEPATH=./cmd/${BINARY_NAME}/main.go

build: tailwind
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ${MAIN_CMD_FILEPATH}
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ${MAIN_CMD_FILEPATH}

air:
	air -c .air.toml

tailwind:
	npm install
	go generate ./assets/tailwind.go

run: build
	./${BINARY_NAME}-darwin

run-linux: build
	./${BINARY_NAME}-linux

clean: 
	go clean
	rm -rf node_modules
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux