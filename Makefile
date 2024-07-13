hot-ui:
	npx vite ui/ --open

all: build

build: 
	go build -o ticketer ./server

run: 
	go run ./server start

gen-clients:
	npx orval --input ./internal/api/api.yaml --output ./ui/src/model

gen: gen-clients
