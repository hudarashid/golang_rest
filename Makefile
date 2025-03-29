# # Build
# .PHONY: build
# build:
# 	@echo "Building local image..."
# 	@docker build -t microservice:local .

# # Test
# .PHONY: test
# test:
# 	@docker compose -f test/docker-compose.yml down -v
# 	@docker compose -f test/docker-compose.yml up --build --abort-on-container-exit --remove-orphans --force-recreate
# 	@docker compose -f test/docker-compose.yml down -v

# Stack
.PHONY:	stop
stop:
	@docker compose -f docker-compose.yml down -v

.PHONY:	prod
prod:
	@docker compose -f docker-compose.yml down -v
	@docker compose -f docker-compose.yml up -d --build

# .PHONY: dev
# dev:
# 	@docker compose -f docker-compose.yml down -v
# 	@docker compose -f docker-compose.yml -f docker-compose.dev.yml up