FROM golang:latest

ARG ENV=PROD
ARG REDIS_URL=redis

RUN mkdir /go/src/shuttleup-kafka
ADD . /go/src/shuttleup-kafka/
WORKDIR /go/src/shuttleup-kafka 
RUN go build -o main .

CMD ["/go/src/shuttleup-kafka/main"]