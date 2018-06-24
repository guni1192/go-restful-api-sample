# GoLang RESTful API

Golang + PostgreSQL + Docker

ORM: gorm


## Deploy

```
$ git clone git@github.com:guni973/go-restful-api-sample
$ docker-compose up -d
```

TODO: APIサーバがPostgreSQLが立ち上がるまで待つようにする

## API Reference

### Hello World

```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080
```

### Index

```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```

### Create

```
$ curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
```

### Update

```
$ curl -XPUT -H 'Content-Type:application/json' http://localhost:8080/users/1 -d '{"name": "koudaiii", "email": "hoge@example.com" }'
```

### Read

```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
```

### DELETE

```
$ curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
```
