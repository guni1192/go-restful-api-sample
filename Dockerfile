FROM golang:1.10.2-alpine

RUN mkdir /app
WORKDIR /app
ADD . /app

EXPOSE 80

CMD ["go", "run", "main.go", "-addr=:80"]
