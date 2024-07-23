all: build

build:
	docker build -t ticketer .

build-local: 
	go build -o ticketer ./server

build-ui: 
	yarn --cwd ui build --outDir ./build/ticketer

run:
	docker run -p 8080:8080 --env-file .env ticketer:latest

# dev
run-local: 
	source .env && go run ./server start

run-worker:
	source .env && go run ./server run-worker

hot-ui:
	npx vite ui/ --open

# codegen

gen-clients:
	npx orval --config ./ui/orval.config.ts

gen-server:
	go generate ./server/... 
	
# && ./internal/api/adjust_generated_schema.py

gen: gen-clients gen-routes gen-db

# db
db-migrate:
	go run ./server migrate
