# Build stage
FROM golang:1.19 AS builder

# Set the working directory

COPY . /app

WORKDIR /app
RUN  go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/main

# Final stage
FROM alpine:3.14 AS final
RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg
# Set the working directory
WORKDIR /app
COPY data.json .
# Copy the binary from the build stage
COPY --from=builder /app/bin/main .

# Expose the port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]

