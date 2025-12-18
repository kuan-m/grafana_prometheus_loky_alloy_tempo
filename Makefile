.PHONY: up

build:
	@cp .env.dev.example .env || true
	@docker compose up -d

up:
	@docker compose up -d
