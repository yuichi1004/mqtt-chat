FROM golang

ADD . /go/src/github.com/yuichi1004/mqtt-chat

RUN go get github.com/yuichi1004/mqtt-chat
RUN go install github.com/yuichi1004/mqtt-chat

WORKDIR /go/src/github.com/yuichi1004/mqtt-chat
ENTRYPOINT /go/bin/mqtt-chat

EXPOSE 8080
