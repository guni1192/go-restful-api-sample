# FROM golang:1.10.2-alpine
FROM golang

# RUN apk add --no-cache git openssl

WORKDIR ${GOPATH}/src/github.com/guni973/go-restful-api-sample
RUN go get -u github.com/golang/dep/cmd/dep
ADD . ./

RUN dep ensure

EXPOSE 8080

CMD ["go", "run", "main.go"]
# CMD tail -f /dev/null
