.PHONY: test
local:
	make orphans
	make test
orphans:
	docker compose -f "docker-compose.yml" down
	docker compose -f "docker-compose-test.yml" down
test:
	docker compose -f "docker-compose.yml" up --build -d
	sleep 10
	docker compose -f "docker-compose-test.yml" up --build --abort-on-container-exit