# Go Agenda
>An Application that provides a way to keep track of your life

## Technologies
The project was developed using Golang,

## Dependencies
* gorilla/mux
> go get -u github.com/gorilla/mux

* logrus
> go get -u github.com/sirupsen/logrus

* gorm
> go get -u github.com/jinzhu/gorm

* Mysql driver
> go get -u github.com/go-sql-driver/mysql

* Mysql dialects
> go get -u github.com/jinzhu/gorm/dialects/mysql

## Endpoints

* check health
> /healthz

* create an item for the todo list
> /item

## Creating database
* Launch a MySQL container  
> docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql

> docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE agenda'

## Testing

* runing agenda
```bash
$ go run agenda.go
```

* checking heath
```bash
$ curl -i localhost:8000/healthz
```

* creating an item
```bash
curl -X POST -d "description=buy apples" localhost:8000/item
```

* updating an item
```bash
curl -X POST -d "done=true" localhost:8000/update/1
```

* deleting an item
```bash
curl -X DELETE localhost:8000/delete/1
```