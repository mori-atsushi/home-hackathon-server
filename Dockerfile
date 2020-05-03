FROM golang:latest

RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

RUN go mod download
RUN go build main.go
CMD ./main
