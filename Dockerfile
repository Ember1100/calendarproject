FROM golang:1.21
  
RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

MAINTAINER "huangqi"

WORKDIR /home/workspace

ADD . /home/workspace

CMD go mod init github.com/calendarproject

CMD go mod tidy

RUN go build main.go

EXPOSE 8016

ENTRYPOINT ["./main"]
