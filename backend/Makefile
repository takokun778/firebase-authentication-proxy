export ENV := dev
export PORT := 8080

init:
	./.develop/init.sh
dev:
	air
run:
	go run main.go
lint:
	golangci-lint run --fix
test:
	go test ./...
key:
	openssl genrsa -out key.pem 2048
di:
	wire
