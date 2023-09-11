.PHONY: up-docker build-docker build-up

up-docker:
	docker-compose up --force-recreate


build-docker:
	docker-compose build

build-up:
	docker-compose up --build