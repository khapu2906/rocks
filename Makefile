
# =========================
# Project Variables
# =========================
include .env
export $(shell sed 's/=.*//' .env)

APP_NAME := user-service
MIGRATE_VERSION := v4.16.2
MIGRATE_BIN := $(GOPATH)/bin/migrate

# =========================
# Install migrate CLI
# =========================
.PHONY: migrate-install
migrate-install:
	@echo "Installing migrate CLI..."
	CGO_ENABLED=1 go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)

# =========================
# Run migrations
# =========================
.PHONY: migrate-up
migrate-up:
	$(MIGRATE_BIN) -path ./migrations -database "${DATABASE_URL}" up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE_BIN) -path ./migrations -database "${DATABASE_URL}" down

.PHONY: migrate-force
migrate-force:
	$(MIGRATE_BIN) -path ./migrations -database "${DATABASE_URL}" force $(VERSION)

.PHONY: migrate-version
migrate-version:
	$(MIGRATE_BIN) -path ./migrations -database "${DATABASE_URL}" version

.PHONY: migrate-create
migrate-create:
	$(MIGRATE_BIN) create -ext sql -dir ./migrations -seq $(name)

# =========================
# App
# =========================
.PHONY: run
run:
	go run cmd/main.go
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

# =========================
# docker-compose setup
# =========================
.PHONY: dev-up
dev-up:
	docker-compose up -d --build

.PHONY: dev-down
dev-down:
	docker-compose down

# =========================
# Generate module
# =========================
.PHONY: generate-module
generate-module:
	scripts/generate_module.sh $(name)


# =========================
# Set up
# =========================
.PHONY: setup
setup:
	go mod download && go mod tidy && make migrate-install

