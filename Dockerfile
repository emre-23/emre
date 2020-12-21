FROM golang:alpine
RUN mkdir /go_challenge
WORKDIR /go_challenge/main
COPY . /go_challenge
RUN go build -o main .
CMD ["/go_challenge/main/main"]