.PHONY: run stop restart build default

run:
	docker compose up -d

stop:
	docker compose down

restart: stop run

build:
	docker compose up --build -d

default:
	echo "default"