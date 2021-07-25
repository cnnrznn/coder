all: build run

build:
	docker compose build

run: build
	docker compose up -d

stop:
	docker compose down