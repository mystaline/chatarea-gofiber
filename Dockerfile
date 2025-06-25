# Build stage
FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# RUN go build -x -o /bin/server ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /bin/server ./cmd/main.go

# Final stage
FROM scratch

WORKDIR /bin

COPY --from=builder /bin/server /bin/

EXPOSE 3000

ENTRYPOINT ["/bin/server"]
