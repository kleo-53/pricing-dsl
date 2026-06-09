FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o pricing-dsl ./pricing-dsl/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/pricing-dsl .
EXPOSE 8080
CMD ["./pricing-dsl"]