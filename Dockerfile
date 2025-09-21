# The base image for the builder stage
FROM golang:1.24-alpine3.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Run commands to tidy and download go modules, then build the application
RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 go build -o appsvc

# Use alpine:latest as the base image for the final stage
FROM alpine:3.22

# Create a non-root user and group
RUN addgroup -S golang && adduser -S gouser -G golang

# Set the working directory in the container
WORKDIR /app

# Copy the built binary from the builder stage to the current directory
COPY --from=datadog/serverless-init:1-alpine /datadog-init /dd/datadog-init
COPY --from=builder /app/appsvc ./

# and change ownership to the non-root user
RUN chown gouser:golang appsvc

# Use the non-root user to run the application
USER gouser

ENTRYPOINT ["/dd/datadog-init"]

# Command to run the application
CMD ["./appsvc"]