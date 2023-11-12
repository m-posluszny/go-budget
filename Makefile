css:
	npx tailwindcss -i ./src/static/tailwind.css -o ./src/static/styling.css --watch;

setup:
	npm install;
	go mod download;


dev:
	go run ./src/main.go

env:
	docker-compose up -d --build;

stop-env:
	docker-compose down;

prod:
	npx tailwindcss -i ./src/static/tailwind.css -o ./src/static/styling.css;
	cd prod && docker-compose up -d --build;

stop-prod:
	cd prod && docker-compose down


test:
	richgo test ./src/... 