FROM golang:1.17

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
