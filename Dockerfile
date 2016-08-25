FROM golang:1.6

ENV GOPATH=/go/src/webapp
RUN go get github.com/lib/pq && go get github.com/gorilla/mux && go get github.com/leekchan/accounting

ADD . /go/src/webapp/

WORKDIR /go/src/webapp

ENTRYPOINT ["go", "run", "src/main/main.go"]
