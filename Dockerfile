FROM golang
RUN mkdir /go-docker
ADD . /go-docker
WORKDIR /go-docker
RUN go build -o main .
CMD ["/go-docker/main"]