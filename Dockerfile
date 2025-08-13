FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM alpine:3.17
COPY --from=builder /app/myapp /myapp
ENTRYPOINT ["/myapp"]
