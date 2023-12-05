# Use the official Golang image for your base image
FROM golang:alpine

WORKDIR /app

# Copy necessary files and folders to the container
COPY . /app

# Build the Go application
RUN go build -o /app/main ./cmd/search-api/

# Expose the port your application runs on
EXPOSE 7000

# Define the command to run your application
CMD ["/app/main"]
