all: build

build: 
	go build -o ticketer ./server

# dev
run-local: 
	source .env && go run ./server start

run-worker:
	source .env && go run ./server run-worker

hot-ui:
	npx vite ui/ --open

# codegen

gen-clients:
	npx orval --config ./internal/api/orval.config.js

gen-routes:
	go generate ./server/...

gen-db:
	go generate ./server/ent

gen: gen-clients gen-routes gen-db

# db
db-migrate:
	go run ./server migrate
