export ENV := dev
export PORT := 8080

init:
	@./.develop/init.sh
dev:
	@air
run:
	@go run main.go
di:
	@wire
lint:
	@golangci-lint run --fix
test:
	@export ENV="test" ;\
	export PRIVATE_KEY=`cat key.pem` ;\
	go test ./...
key:
	@openssl genrsa -out key.pem 2048 && PRIVATE_KEY=`cat key.pem` && echo $${PRIVATE_KEY}
