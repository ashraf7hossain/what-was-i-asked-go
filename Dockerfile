FROM golang:latest

WORKDIR /go/src/app

# Install CompileDaemon
# RUN go install github.com/githubnemo/CompileDaemon@latest

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Run migration
RUN go run migrate/migrate.go

# Build the application
RUN go build -o rest-in-go .

# Command to run CompileDaemon with proper arguments
CMD ["CompileDaemon", "--build=go build -o rest-in-go .", "--command=./rest-in-go"]