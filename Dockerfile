# Server builder
FROM golang:1.22-alpine AS server_builder

WORKDIR /app
COPY . .
RUN cd server && go mod download
RUN go build -o ticketer ./server

# UI Builder







# Start a new stage from scratch
FROM alpine:latest  
WORKDIR /app/
COPY --from=server_builder /app/ticketer .

# Copy any additional necessary files (like .env if needed)
# COPY .env .

EXPOSE 8080
CMD ["./ticketer", "start"]