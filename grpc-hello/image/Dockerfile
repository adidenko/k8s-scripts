FROM golang:1.22-alpine3.19

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN apk add git \
    && git clone -b v1.67.0 --depth 1 https://github.com/grpc/grpc-go \
    && cd grpc-go/examples/helloworld

ADD ./greeter_client/main.go /go/grpc-go/examples/helloworld/greeter_client

RUN cd grpc-go/examples/helloworld \
    && go install google.golang.org/grpc/examples/helloworld/greeter_server \
    && go install google.golang.org/grpc/examples/helloworld/greeter_client

ENTRYPOINT [ "/go/bin/greeter_server" ]
