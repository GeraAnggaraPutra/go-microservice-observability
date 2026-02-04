FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app_binary .

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache tzdata

ENV TZ=Asia/Jakarta

COPY --from=builder /build/app_binary /app/app_binary

CMD ["/app/app_binary"]