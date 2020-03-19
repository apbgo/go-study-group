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

chapter6_exec_sample1: chapter6_migrate
	go run chapter6/sample1/main.go

chapter6_exec_sample2: chapter6_migrate
	go run chapter6/sample2/main.go

chapter6_exec_sample3: chapter6_migrate
	go run chapter6/sample3/main.go

chapter6_exec_sample4: chapter6_migrate
	go run chapter6/sample4/main.go

chapter6_exec_sample5: chapter6_migrate
	go run chapter6/sample5/main.go

chapter6_exec_sample6: chapter6_migrate
	go run chapter6/sample6/main.go

chapter6_exec_sample7: chapter6_migrate
	go run chapter6/sample7/main.go

chapter6_exec_sample8: chapter6_migrate
	go run chapter6/sample8/main.go