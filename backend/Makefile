export ENV := dev
export PORT := 8080

init:
	./.develop/init.sh
dev:
	air
run:
	go run main.go
check:
	go vet ./...
	staticcheck ./...
	gosec ./...
test:
	go test ./...
key:
	openssl genrsa -out key.pem 2048
