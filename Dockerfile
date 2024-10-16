# Build stage
FROM golang:1.19 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /app/bin/main ./cmd/main.go

# Final stage
FROM alpine:3.14
RUN apk update && apk upgrade && apk add --no-cache ffmpeg
WORKDIR /app
COPY --from=builder /app/bin/main .
COPY --from=builder /app/data.json .
COPY .env .

EXPOSE 8080
CMD ["./main"]
