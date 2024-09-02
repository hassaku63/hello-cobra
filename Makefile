BIN_DIR := bin

.PHONY: build
build:
	go build -o ${BIN_DIR}/main ./main

.PHONY: clean
clean:
	rm -rf ${BIN_DIR}/
