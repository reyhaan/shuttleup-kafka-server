FROM golang:latest

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
# RUN apt-get update; \
#     apt-get install -y ca-certificates
# RUN go get
RUN go build -o main .

CMD ["/app/main"]