BINARY_NAME=simpbb
CMD_PATH=./cmd/${BINARY_NAME}/${BINARY_NAME}.go

build: generate-assets
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ${CMD_PATH}
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${CMD_PATH}

air:
	air -c .air.toml

generate-assets:
	npm install
	go generate ./assets/embed.go

start: build
	./${BINARY_NAME}-darwin start

start-linux: build
	./${BINARY_NAME}-linux start

clean: 
	go clean
	rm -rf node_modules
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux