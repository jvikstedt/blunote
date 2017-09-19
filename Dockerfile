FROM golang:1.9

run go get -u github.com/golang/dep/cmd/dep
run go get -u github.com/codegangsta/gin

run mkdir -p /go/src/github.com/jvikstedt/blunote
WORKDIR /go/src/github.com/jvikstedt/blunote

COPY Gopkg.lock Gopkg.toml /go/src/github.com/jvikstedt/blunote/
run dep ensure -vendor-only

copy . /go/src/github.com/jvikstedt/blunote
