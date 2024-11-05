FROM golang:1.22@sha256:0ca97f4ab335f4b284a5b8190980c7cdc21d320d529f2b643e8a8733a69bfb6b

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
