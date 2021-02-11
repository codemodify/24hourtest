FROM golang:1.15

WORKDIR $GOPATH/src/github.com/codemodify/24hourtest/server

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...


EXPOSE 64000

CMD ["server"]