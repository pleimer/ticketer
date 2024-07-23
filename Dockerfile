# Server builder
FROM golang:1.22-alpine AS server_builder

WORKDIR /app
COPY . .
RUN cd server && go mod download
RUN go build -o ticketer ./server

# UI Builder
FROM node:22 AS ui_builder
WORKDIR /app

COPY ui ./ui
RUN yarn --cwd ui install
RUN yarn --cwd ui build --outDir ./build/ticketer

# Start a new stage from scratch
FROM alpine:latest  
WORKDIR /app/
COPY --from=server_builder /app/ticketer ./server/ticketer
COPY --from=ui_builder /app/ui/build/ticketer ./ui/build/ticketer

CMD ./server/ticketer start