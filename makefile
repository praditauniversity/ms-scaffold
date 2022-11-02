.PHONY: test
test:
	docker compose down
	docker compose up --abort-on-container-exit