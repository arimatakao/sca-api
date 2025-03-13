run:
	go run main.go -config ./config.yaml
build:
	mkdir bin
	go build -o ./bin/sca-api main.go
image:
	docker build --no-cache -t sca-api:latest .
test:
	go test -v ./...
clear:
	rm -rf ./bin