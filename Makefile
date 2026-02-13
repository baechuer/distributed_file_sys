build:
	@go build -o bin/fs main.go

run: build
	@./bin/fs

test:
	@go test ./... -v

clean:
	@rm -rf bin

