hot-ui:
	npx vite ui/ --open

all: build

build: 
	go build -o ticketer ./server

run: 
	go run ./server start

gen-clients:
	npx orval --config ./internal/api/orval.config.js

gen: gen-clients
