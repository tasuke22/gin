.PHONY: help

help:
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

cp-env:
	cp .env.example .env
up:
	docker compose up -d
db:
	docker compose exec -it db mysql -u myuser -pmypassword -D mydatabase
logs:
	docker compose logs -f
mup:
	migrate -database "mysql://myuser:mypassword@tcp(localhost:3306)/mydatabase?query" -path ./migrations up
mdown:
	migrate -database "mysql://myuser:mypassword@tcp(localhost:3306)/mydatabase?query" -path ./migrations down
