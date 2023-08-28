.phony:

run:
	cd webapp/golang/ && go run app.go

env:
	export $$(cat .env)

up:
	docker compose up -d
