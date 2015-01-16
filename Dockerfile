FROM golang:1.4
MAINTAINER Shad Beard

ADD . /go/src/github.com/Dahs81/sq
RUN go install github.com/Dahs81/sq

ENTRYPOINT /go/bin/sq
