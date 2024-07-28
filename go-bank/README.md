## Go-bank

JSON API server in Golang that runs on port 3000

```
GET http://localhost:3000/account [Get all accounts]

POST http://localhost:3000/account [Create new account]
```

### How to run the application

Start server: `make run`


### Packages installed

```
go get github.com/gorilla/mux
go get github.com/lib/pq
```