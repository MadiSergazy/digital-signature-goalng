.PHONY: up-docker build-docker build-up

up-docker:
	docker compose up


build-docker:
	docker compose build

build-up:
	docker compose up --build