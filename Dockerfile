# Stage 1 - Builder
FROM golang:1.23-alpine AS builder

# Set environment variables for Go
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

RUN apk add --no-cache git
RUN go install github.com/google/wire/cmd/wire@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN wire ./cmd/wire.go
RUN go build -o gymtrack main.go

# Stage 2 - Final Image
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/gymtrack .
COPY .env* ./
EXPOSE 5555
CMD ["./gymtrack", "serve"]
