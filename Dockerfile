FROM golang:1.6

RUN go get github.com/lib/pq && go get github.com/gorilla/mux && go get github.com/leekchan/accounting

RUN mkdir /go/src/webapp
ADD . /go/src/webapp/

WORKDIR /go/src/webapp
RUN sh env.sh
