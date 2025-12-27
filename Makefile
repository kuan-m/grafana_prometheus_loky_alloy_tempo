-include .env

ENV = .env
DUMP_FILE = infra/dumps/dump.sql
MYSQL_DATABASE ?= prometheus
DB_ROOT_PASSWORD ?= root
DB_USERNAME = root

.PHONY: up down gen-logs import-db

build:
	@cp .env.example .env || true
	@docker compose up -d

up:
	@docker compose up -d

down:
	@docker compose down

gen-logs:
	@echo "Generating logs..."
	@docker compose run --rm log-generator

# =========================================
# Database import
# =========================================
import-db:
	@if [ ! -f "$(DUMP_FILE)" ]; then \
		echo "❌ Dump file not found: $(DUMP_FILE)"; \
		exit 1; \
	fi
	@echo "======================================="
	@echo "📥 Importing database..."
	@echo "File: $(DUMP_FILE)"
	@echo "Database: $(MYSQL_DATABASE)"
	@echo "======================================="
	@echo "1. Dropping existing database..."
	@docker exec mysql \
		mysql -u$(DB_USERNAME) -p$(DB_ROOT_PASSWORD) \
		-e "DROP DATABASE IF EXISTS $(MYSQL_DATABASE);" || true
	@echo "2. Creating new database..."
	@docker exec mysql \
		mysql -u$(DB_USERNAME) -p$(DB_ROOT_PASSWORD) \
		-e "CREATE DATABASE IF NOT EXISTS $(MYSQL_DATABASE) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
	@echo "3. Importing data..."
	@docker exec -i mysql \
		mysql -u$(DB_USERNAME) -p$(DB_ROOT_PASSWORD) $(MYSQL_DATABASE) < $(DUMP_FILE)
	@echo "======================================="
	@echo "✅ Database successfully imported!"
	@echo "======================================="
