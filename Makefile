hot-ui:
	npx vite ui/ --open

# Default target is "build"
all: build

# Build target compiles your Go source files into a single executable binary
build: main.go
	go build -o ticketer server/cmd/main.go

# Run target builds and then runs the binary
run: 
	cd server && go run main.go