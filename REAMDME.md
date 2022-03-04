# set up



RUN cd -p calc/design && \
    cd calc && \
    go mod init calc && \
    go get -u goa.design/goa/v3 && \
    go get -u goa.design/goa/v3/... && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get -d google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest