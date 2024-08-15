# Use the official Go image from the Docker Hub
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /TevianTestTask

# Copy the Go modules manifests
COPY go.mod go.sum  ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN ls

# Build the Go app
RUN go build -o tevian_test_task ./cmd/tevian_test_task

# Use a smaller base image for the runtime
FROM alpine:latest

# Install any required dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/TevianTestTask

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /TevianTestTask/tevian_test_task .
COPY --from=builder /TevianTestTask/internal/api/docs ./internal/api/docs

# Expose port 3000 to the outside world
EXPOSE 3000

ENV FACECLOUD_EMAIL=kiselyovvld@mail.ru
ENV FACECLOUD_PASSWORD=Lesstenpound@159951
ENV FACECLOUD_API_URL=https://backend.facecloud.tevian.ru/api

ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=database

# Command to run the executable
CMD ["./tevian_test_task"]