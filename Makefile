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
	npx orval --config ./ui/orval.config.ts

gen-server:
	go generate ./server/... && ./internal/api/adjust_generated_schema.py

# gen-db:
# 	go generate ./server/ent 

gen: gen-clients gen-routes gen-db

# db
db-migrate:
	go run ./server migrate
