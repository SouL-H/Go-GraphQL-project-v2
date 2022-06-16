
# GraphQL 

```
go run github.com/99designs/gqlgen init

```

## Applied after schema change.
```
go run github.com/99designs/gqlgen generate
```

## Dockerize Command
```
docker compose down && docker compose up --build

default uri : http://localhost:8080/
```

## Query Example


devtopics


<p>
    <img src="./img/query.png"  style="width:600px;" alt="Observer">

</p>

### Different Query

<p>
<img src="./img/query2.png"  style="width:600px;" alt="Observer">

</p>

## Create Todo

<p>
<img src="./img/createTodo.png"  style="width:600px;" alt="Observer">

</p>

## To Reduce Docker File size by %60
```
./Dockerfile

FROM golang 

Replace with " FROM golang:1.16-alpine "

Previous size 860MB
Current size is 330MB
``` 
#### Thank you devtopics 