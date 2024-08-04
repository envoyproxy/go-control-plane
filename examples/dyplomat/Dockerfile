FROM golang:1.21@sha256:672a2286da3ee7a854c3e0a56e0838918d0dbb1c18652992930293312de898a6

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
