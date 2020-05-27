FROM golang:alpine

WORKDIR $GOPATH/src/mylekkepackage/mylekkeapp/
ADD ./ .
RUN go build -o /go/code-edu-calc