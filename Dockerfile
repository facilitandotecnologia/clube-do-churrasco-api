# Use the official Go image as a base
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the Go source code into the container
COPY . /app

# Set the environment variables
ENV GOOS linux
ENV GOARCH amd64

# Build the Go application
RUN go build -o main ./cmd/server

# Expose the port that the application will use
EXPOSE 8080

# Run the command to start the application when the container is launched
CMD ["./main"]
