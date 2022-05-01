FROM golang:1.17

RUN apt-get update && \
    apt-get install unzip

RUN apt-get install -y protobuf-compiler


WORKDIR /go/src/work

ENV GO111MODULE on

# install go tools
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls \
  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \