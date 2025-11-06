FROM golang:1.24@sha256:8b788ecf5fddb91cd04e532205d6411de68a08b510885bb8fb1c93f54c03f737

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
