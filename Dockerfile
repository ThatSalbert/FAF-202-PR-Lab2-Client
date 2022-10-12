FROM golang:alpine

RUN mkdir /client

WORKDIR /client

COPY . .

RUN go build -o /go/bin/main

EXPOSE 6000

ENTRYPOINT ["/go/bin/main"]

