BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} cmd/main.go

run:
	go build -o ${BINARY_NAME} cmd/main.go
	./${BINARY_NAME}

test:
	go test ./...
clean:
	go clean
	rm ${BINARY_NAME}