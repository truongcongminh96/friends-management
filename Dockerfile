FROM golang:1.16 as builder

COPY go.mod go.sum /home/sites/friends-management/

WORKDIR /home/sites/friends-management/

RUN go mod download

COPY . /home/sites/friends-management/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/friends-management home/sites/friends-management

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /home/sites/friends-management/build/friends-management /usr/bin/friends-management
EXPOSE 8080 8081
ENTRYPOINT ["/usr/bin/friends-management"]