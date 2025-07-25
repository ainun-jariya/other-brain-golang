FROM golang:1.24.5 AS builder
WORKDIR /app
COPY . .
ENV port=8080
RUN go build -o server main.go

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/server .
ENV port=8080
EXPOSE 8080
CMD ["/app/server"]
