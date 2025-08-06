FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -a -o server ./cmd/server

FROM scratch

WORKDIR /

COPY --from=builder /app/server ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY static ./static
COPY templates ./templates
COPY content ./content

EXPOSE 8080

CMD ["./server"]
