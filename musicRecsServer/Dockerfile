# Start from the official Golang image
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules to download dependencies
COPY go.mod ./
RUN go mod download

# Copy all Go files
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -o ./musicRecsServer

# Expose the server port
EXPOSE 8080

# Run the server
CMD ["./musicRecsServer"]