include .env
export

.PHONY: up
up:
	@docker compose --project-name function --file ./.docker/compose.yaml up -d

.PHONY: down
down:
	@docker compose --project-name function down --volumes

.PHONY: reload
reload:
	@touch local/main.go
