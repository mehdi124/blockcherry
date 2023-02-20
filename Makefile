build:
	go build -o ./bin/blockcherry
run: build
	./bin/blockcherry
test:
	go test -v ./...
