# 24hourtest


# 1 - clone the repo
```bash
export GO111MODULE=on
export GOPATH="`cd ~/go;pwd`"
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

go get github.com/codemodify/24hourtest
```

# 2 - run Docker image
```bash
cd $GOPATH/src/github.com/codemodify/24hourtest/demo-workspace
sudo docker build -t nicolae24hourtest .
sudo docker run nicolae24hourtest
```

# 3 - consume GRPC service from Docker image
```bash
cd $GOPATH/src/github.com/codemodify/24hourtest/demo-workspace/client
go build
./client

# server will say "DEBUG: sending hello world"
# client will say "DEBUG: got: hello world"

```


# 4 - Run/Debug server/client from VSCode
- VSCode -> open folder -> open "24hourtest" folder
- in VSCode terminal `cd server && go run main.go`
- open client/main.go in VSCode editor, put a breakpoint in main function
- hit Debug/F5
