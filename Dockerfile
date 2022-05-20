FROM golang:1.18.1-bullseye

ENV AWS_SDK_LOAD_CONFIG=1

WORKDIR /usr/local/app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/ramya-rao-a/go-outline@latest && \
    go install -v github.com/cweill/gotests/gotests@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v github.com/golang/mock/mockgen@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    \
    go mod tidy
