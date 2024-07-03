run dev:
	go run main.go -mode debug

run test:
	go run main.go -mode test

build release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/release/

build test:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/test/

build dev:
	go build -o target/dev/
