lint:
	@echo "Running revive..."
	@revive -config revive.toml -formatter stylish ./...

	@echo "Running errcheck..."
	@errcheck ./...

build:
	go build -o CsvManipulator ./...

all: lint build

.PHONY: lint build all
