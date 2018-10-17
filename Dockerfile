FROM golang:1.11-alpine as builder
# update and upgrade apk
RUN apk update; apk upgrade
# install git
RUN apk add --no-cache git
# install dep
RUN go get github.com/golang/dep/cmd/dep
# create a working directory
WORKDIR /go/src/github.com/phaicom/golang-docker-example
# add Gopkg.toml and Gopkg.lock
COPY Gopkg.toml Gopkg.lock ./
# install packages
RUN dep ensure --vendor-only
# add source code
COPY . ./
# build app.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./app.go

# use a minimal alpine image
FROM alpine:3.8
# set working directory
WORKDIR /go/src
# copy the src file from builder
COPY --from=builder /go/src/github.com/phaicom/golang-docker-example/ ./
# run the binary
CMD ["./app"]