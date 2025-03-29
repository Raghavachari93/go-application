# Use the official Golang image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code
COPY . .

# Build the Go application
RUN go build -o go-app

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/app/go-app"]
