# Use golang:alpine as the base image for the builder stage
FROM golang:alpine as builder

# Install git, necessary for fetching Go dependencies
RUN apk add --no-cache git

# Create a directory for the application build
RUN mkdir /build 

# Copy the local code to the container
ADD . /build/

# Set the working directory for the build commands
WORKDIR /build 

# Build the Go application, disabling CGO for a fully static build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine

# Create a non-root user and group for running the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the user and group IDs for security
USER appuser:appgroup

# Create a directory for the application
RUN mkdir /app

# Set the working directory to /app
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /build/main /app/

# Expose port 3000 for the application
EXPOSE 3000

# Command to run the application
CMD ["./main"]
