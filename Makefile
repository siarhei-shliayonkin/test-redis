RUN	= go run
BUILD	= go build
BIN_DIR	= ./bin

REDIS_PROJECT_DIR	= ./cmd
REDIS_PROJECT_NAME	= redis

.PHONY: run build bindir
run:
	@$(RUN) $(REDIS_PROJECT_DIR)
build: bindir
	@$(BUILD) -o $(BIN_DIR)/$(REDIS_PROJECT_NAME) $(REDIS_PROJECT_DIR)/main.go
bindir:
	@if [ ! -d $(BIN_DIR) ]; then echo "binary dir does not exist, creating.."; mkdir $(BIN_DIR); fi
