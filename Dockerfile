# Use a minimal base image
FROM golang:1.21 as builder

# Set the working directory
WORKDIR /app

# Copy only the necessary files to minimize the build context
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kandji-prometheus-exporter .

# Use a minimal base image for the final image
FROM scratch

# Set the working directory
WORKDIR /app

# Copy only the binary from the previous stage
COPY --from=builder /app/kandji-prometheus-exporter .

# Expose the port on which the exporter will run
EXPOSE 8080

# Run the exporter
CMD ["./kandji-prometheus-exporter"]
