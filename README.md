# golang-docker-example
[![Build Status](https://travis-ci.org/Phaicom/golang-docker-example.svg?branch=master)](https://travis-ci.org/Phaicom/golang-docker-example)

## Prerequisites

* [Docker](https://www.docker.com/)
* [Go](https://golang.org)
* [Dep](https://github.com/golang/dep) - Go dependency management tool 

## Installing

Create a Docker image

```
$ docker build -t golang-docker-example:1.0 .
```

Start a container

```
$ docker run -p 8080:8080 golang-docker-example:1.0
```

Or if you have go installed

```
$ dep ensure
$ go run app.go
```

## Author

By [Reawpai Chunsoi](https://github.com/phaicom/)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details