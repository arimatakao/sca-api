run:
	go run main.go
build:
	mkdir bin
	go build -o ./bin/sca-api main.go
test:
	go test -v ./...
clear:
	rm -rf ./bin