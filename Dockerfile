# Server builder
FROM golang:1.22-alpine AS server_builder

WORKDIR /app
COPY . .
RUN cd server && go mod download
# codegen
RUN go generate ./server/... 
RUN go build -o ticketer ./server

# UI Builder
FROM node:22 AS ui_builder
WORKDIR /app

COPY ui ./ui
COPY internal ./internal
RUN yarn --cwd ui install
RUN npx orval --config ./ui/orval.config.ts
RUN yarn --cwd ui build --outDir ./build/ticketer

# Build final image
FROM alpine:latest  
WORKDIR /app/
COPY --from=server_builder /app/ticketer ./server/ticketer
COPY --from=ui_builder /app/ui/build/ticketer ./ui/build/ticketer

CMD ./server/ticketer start