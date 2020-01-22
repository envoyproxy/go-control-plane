FROM golang:1.13.6-stretch

WORKDIR /go-control-plane

# Fetch and preserve module dependencies
ENV GOPROXY=https://proxy.golang.org
COPY go.mod ./
RUN go mod download

# Fetch protoc modules
RUN go get github.com/envoyproxy/protoc-gen-validate@d6164de4910977d3c3c8dbd9299b5064ea9e7c2b

# Install protoc 3.6.1
RUN apt-get update && apt-get install unzip
RUN wget -O /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip  \
  && unzip /tmp/protoc.zip -d /usr/local/ \
  && rm /tmp/protoc.zip

# circle does not like files in the checkout directory
RUN rm go.mod
RUN rm go.sum

# add envoy
COPY --from=envoyproxy/envoy:v1.13.0 /usr/local/bin/envoy /usr/local/bin/envoy
