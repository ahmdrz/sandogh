FROM golang:latest

RUN mkdir -p $GOPATH/src/github.com/ahmdrz/sandogh
COPY . $GOPATH/src/github.com/ahmdrz/sandogh
WORKDIR $GOPATH/src/github.com/ahmdrz/sandogh

RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -race -i -o storage

EXPOSE 8080

CMD ["./storage"]