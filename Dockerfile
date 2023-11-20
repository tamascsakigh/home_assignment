FROM golang:1.20 as gobuilder

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin"
RUN chmod -R 777 "$GOPATH"

WORKDIR /go/src/github.com/tamascsakigh/home_assignment

COPY . .

RUN go mod tidy
RUN go mod vendor

RUN go build -o main main.go


FROM debian:bookworm

RUN cp /usr/share/zoneinfo/Europe/Budapest /etc/localtime
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /

COPY --from=gobuilder /go/src/github.com/tamascsakigh/home_assignment/main .

ENV CONFIG_FILE="./.env"

CMD ["./main"]