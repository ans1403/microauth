FROM golang:1.18.1-bullseye

WORKDIR /usr/local/golang-devcontainer

RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/ramya-rao-a/go-outline@latest && \
    go install -v github.com/cweill/gotests/gotests@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v github.com/golang/mock/mockgen@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest
