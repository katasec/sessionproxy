FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o sessionproxy

FROM gcr.io/distroless/base
COPY --from=builder /app/sessionproxy /

CMD ["/sessionproxy"]