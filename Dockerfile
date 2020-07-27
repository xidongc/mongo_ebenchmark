FROM golang:1.14-alpine3.11 as builder
COPY . /go/src/github.com/xidongc/mongodb_ebenchmark
WORKDIR /go/src/github.com/xidongc/mongodb_ebenchmark
RUN go build -o /ebenchmark github.com/xidongc/mongodb_ebenchmark/cmd

FROM alpine:3.11
ENV GOTRACEBACK=single
CMD ["./ebenchmark"]
RUN mkdir /protofile
COPY pkg/proxy/rpc.protoset /protofile
ENV PROTOSET_FILE=/protofile/rpc.protoset
COPY --from=builder /ebenchmark .
