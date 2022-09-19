FROM golang:alpine

RUN mkdir /dining_hall

WORKDIR /dining_hall

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8080

ENTRYPOINT ["/go/bin/main"]

