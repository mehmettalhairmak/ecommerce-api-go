# Go base image
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# --- final image ---
FROM debian:bullseye-slim

WORKDIR /root/

# Sadece binary'yi al
COPY --from=builder /app/server .

# Port aç
EXPOSE 4242

# Uygulamayı başlat
CMD ["./server"]