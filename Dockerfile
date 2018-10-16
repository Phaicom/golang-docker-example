FROM golang:1.11-alpine
# update and upgrade apk
RUN apk update; apk upgrade
# install git
RUN apk add git
# install dep
RUN go get github.com/golang/dep/cmd/dep
# create a working directory
WORKDIR $GOPATH/src/github.com/phaicom/golang-docker-example
# add Gopkg.toml and Gopkg.lock
COPY Gopkg.toml Gopkg.lock ./
# install packages
RUN dep ensure --vendor-only
# add source code
COPY . ./
# build main.go
RUN go build ./app.go
# run the binary
CMD ["./app"]