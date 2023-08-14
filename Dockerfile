FROM golang:1.21.0-alpine as builder

WORKDIR /app

COPY ./ ./
RUN go mod download

ENV GOOS=linux
ENV CGO_ENABLED=0

RUN go build \
    -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

ENV TZ="UTC"

COPY --from=builder /app/main .

CMD ["/app/main"]