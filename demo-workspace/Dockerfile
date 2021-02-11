FROM golang:alpine
RUN apk update
RUN apk add git

ENV GO111MODULE=on \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	GOPATH="/go" \
	GOBIN=$GOPATH/bin \
	PATH=$PATH:$GOBIN

# RUN go get github.com/codemodify/24hourtest; cd $GOPATH/src/github.com/codemodify/24hourtest; ./compile-and-run.sh
RUN mkdir -p $GOPATH/src/github.com/codemodify
RUN cd $GOPATH/src/github.com/codemodify
RUN git clone https://github.com/codemodify/24hourtest.git
RUN ./compile-and-run.sh
