# set up



RUN cd -p calc/design && \
    cd calc && \
    go mod init calc && \
    go get -u goa.design/goa/v3 && \
    go get -u goa.design/goa/v3/... && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# proto用ドキュメント生成

```
protoc --doc_out=./gen/grpc --doc_opt=markdown,index.md gen/grpc/calc/pb/*.proto

```

[pseudomuto/protoc-gen-doc: Documentation generator plugin for Google Protocol Buffers](https://github.com/pseudomuto/protoc-gen-doc)