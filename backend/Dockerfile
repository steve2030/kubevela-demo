# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod .

# Download dependencies (if any)
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o library-backend .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./library-backend"]