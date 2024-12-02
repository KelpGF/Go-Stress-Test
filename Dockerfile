FROM golang:1.23 AS builder

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o stress .

FROM scratch
COPY --from=builder /app/stress .
ENTRYPOINT ["./stress"]
