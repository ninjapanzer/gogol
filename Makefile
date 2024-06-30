PHONY: build-all build-tradgol build-parallelgol

default: build-all

build-all: build-tradgol build-parallelgol

build-tradgol:
	@echo "Building traditional gol..."
	@go vet ./cmd/tradgol
	@go build -o gol ./cmd/tradgol

build-parallelgol:
	@echo "Building parallel gol..."
	@go vet ./cmd/parallelgol
	@go build -o parallelgol ./cmd/parallelgol
