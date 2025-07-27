FROM golang:1.24.5 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o server main.go

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/server .
ENV PORT=8080
EXPOSE 8080

CMD ["/app/server"]
