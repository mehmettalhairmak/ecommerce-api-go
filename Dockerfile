# Go base image
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# --- final image ---
FROM debian:bookworm-slim

WORKDIR /root/


COPY --from=builder /app/server .

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*


EXPOSE 4242

# Uygulamayı başlat
CMD ["./server"]