FROM golang:1.16-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app shortener
RUN mkdir /app
RUN chown shortener /app
RUN chmod 0744 /app
USER shortener
COPY --from=builder /build/golang-url-shortener /app/
COPY templates/ /app/templates
WORKDIR /app
CMD ["./golang-url-shortener"]