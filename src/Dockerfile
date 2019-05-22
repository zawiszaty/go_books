FROM golang:1.12.4
RUN go get github.com/codegangsta/gin

WORKDIR $GOPATH/src/go_books/src

COPY . .

RUN go get
RUN go get github.com/pilu/fresh
