# Lint, Format
lint:
	go fmt ./...
	goimports -w ./
	go vet ./...

# Test
test:
	go test -v -race ./...

chapter6_migrate:
	mysql -h127.0.0.1 -P 5446 -uroot < chapter6/migraiton.sql