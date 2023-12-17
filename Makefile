.PHONY:unit-test
unit-test:
	@echo "running unit tests"
	go test -v `go list ./... | grep -v test/`

.PHONY:coverage
coverage:
	@echo "coverage report"
	go test -cover -v -timeout 60s -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY:clean
clean:
	@echo "cleaning the goriot package"
	rm -rf coverage.out coverage.html