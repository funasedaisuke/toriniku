FROM golang:1.13

RUN apt-get update -qq \
    && apt-get install -y default-mysql-client vim \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/jinzhu/gorm/dialects/mysql

WORKDIR /go/src/toriniku

ADD . .

RUN go build