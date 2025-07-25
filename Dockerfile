FROM golang:1.24.5 AS builder
WORKDIR /app
COPY . .
RUN go build -o server main.go

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["PORT=8080 /app/server"]
