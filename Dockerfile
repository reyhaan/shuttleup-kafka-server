FROM golang:latest

ARG ENV=PROD
ARG REDIS_URL=redis

ENV ENV=$ENV
ENV REDIS_URL=$REDIS_URL

RUN mkdir /go/src/shuttleup-kafka
ADD . /go/src/shuttleup-kafka/
WORKDIR /go/src/shuttleup-kafka 
RUN go build -o main .

CMD ["/go/src/shuttleup-kafka/main"]