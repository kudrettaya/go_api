# ---- Build stage ----
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

# Copy dependency files first so the download layer is cached
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary from the cmd package
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd

# ---- Run stage ----
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /server .

EXPOSE 8000

CMD ["./server"]