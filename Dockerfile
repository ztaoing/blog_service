FROM golang:latest as builder
RUN mkdir -p /go/src/blog_service
WORKDIR /go/src/blog_service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o blog .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/blog_service/blog .
CMD ["./blog"]