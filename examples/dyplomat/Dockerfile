FROM golang:1.21

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
