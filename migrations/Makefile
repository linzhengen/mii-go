.phony: migrate-up migrate-down migrate-create
#DB_USR := $(shell grep DB_USER .env | cut -d '=' -f2)
#DB_PASSWORD := $(shell grep DB_PASSWORD .env | cut -d '=' -f2)
#DB_HOST := $(shell grep DB_HOST .env | cut -d '=' -f2)
#DB_PORT := $(shell grep DB_PORT .env | cut -d '=' -f2)
#DB_NAME := $(shell grep DB_NAME .env | cut -d '=' -f2)
#
migrate-create:
	@read -p "Enter migration name: " name; \
	if [ -z "$$name" ]; then \
		echo "Migration name is required"; \
		exit 1; \
	fi; \
	migrate create -seq=true --ext=sql -dir=./mysql $$name

#migrate-up:
#	migrate -path ./mysql -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose up
#
#migrate-down:
#	migrate -path ./mysql -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose down