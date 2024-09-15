run: 
	go run cmd/main.go

# Run tests
test-unit:
	go test -v -cover -coverprofile=c.out -coverpkg=./... test/*_unit_*.go