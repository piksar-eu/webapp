WEBSITE_DIR := ./apps/website
CORE_DIR := ./apps/core
CORE_MAIN_FILE := cmd/main.go
OUTPUT_DIR := dist
OUTPUT_BINARY := $(OUTPUT_DIR)/core

.PHONY: clean

all: build_website build_core

run_website:
	cd $(WEBSITE_DIR) && \
		npm run dev

run_core:
	cd $(CORE_DIR) && \
		go run cmd/main.go

build_website: $(WEBSITE_DIR)/**/*
	cd $(WEBSITE_DIR) && \
		npm run build

build_core: $(CORE_DIR)/**/*.go
	mkdir -p $(OUTPUT_DIR)
	cd $(CORE_DIR) && \
		go build -o ../../$(OUTPUT_BINARY) $(CORE_MAIN_FILE)

clean:
	rm -rf $(OUTPUT_DIR) && \
	rm -rf $(WEBSITE_DIR)/dist
