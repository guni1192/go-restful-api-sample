FROM golang:1.10.2-alpine

RUN mkdir /app
WORKDIR /app
ADD . /app

EXPOSE 8080

CMD ["go", "run", "main.go"]
