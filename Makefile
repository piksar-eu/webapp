WEBSITE_DIR := ./apps/website
CORE_DIR := ./apps/core
CORE_MAIN_FILE := cmd/main.go
OUTPUT_DIR := dist

.PHONY: clean

all: core_build

website_install: $(WEBSITE_DIR)/package.json
	cd $(WEBSITE_DIR) && \
		npm install

website_run: website_install
	cd $(WEBSITE_DIR) && \
		npm run dev

website_build: $(WEBSITE_DIR)/**/* website_install
	cd $(WEBSITE_DIR) && \
		$(shell sed -n 's/^\(VITE_[^=]*\)=.*/\1=\1/p' .env_example) \
		WEBSITE_DIST_DIR=apps/core/pkg/web/static/website \
		npm run build

core_run:
	cd $(CORE_DIR) && \
		ENV_PATH=../../.env go run cmd/main.go

core_build: $(CORE_DIR)/**/*.go website_build
	mkdir -p $(OUTPUT_DIR)
	cd $(CORE_DIR) && \
		go build -o ../../$(OUTPUT_DIR)/core $(CORE_MAIN_FILE)

clean:
	rm -rf $(OUTPUT_DIR) && \
	rm -rf $(WEBSITE_DIR)/dist
