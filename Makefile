
benchmark:
	@echo "Benchmarking..."
	@go test -bench=. -benchmem -benchtime=3s -cpu=10

test:
	@echo "Testing..."
	@go test -v -cover