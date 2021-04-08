FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
COPY .env .
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
