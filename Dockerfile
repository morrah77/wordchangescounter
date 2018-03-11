FROM golang:1.9.1

WORKDIR /go/src/github.com/morrah77/wordchangescounter
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure
RUN ./control.sh build

CMD `./control.sh run`
