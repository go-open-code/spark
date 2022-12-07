PKG := github.com/go-open-code/spark
BIN_DIR ?= ./bin/

spark:
	go build -o $(BIN_DIR) ./cmd/spark