# productLab
 Lab for testing, sql, golang etc ..

- setting project

    export APP_DB_USERNAME=postgres
    export APP_DB_PASSWORD=
    export APP_DB_NAME=postgres

    docker run -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres

    go mod tidy && go mod vendor

    go test ./...

    go run main.go

- generate mocks: 
    mockgen -destination=mocks/mock_productstore.go -package=mocks lab/productLab/internal/usecase ProductStore
    mockgen -destination=mocks/mock_productus.go -package=mocks lab/productLab/internal/transport ProductUC

## Interacting
- create product:
curl -H "Content-Type: application/json" -X POST -d '{"id":1, "name":"jabon", "price":3.12}' http://localhost:8010/products

- get product
curl -X GET http://localhost:8010/products/1

- update product
curl -H "Content-Type: application/json" -X PUT -d '{"name":"te", "price":30.12}' http://localhost:8010/products/1

- list products
curl -X GET http://localhost:8010/products

- delete product
curl -X DELETE http://localhost:8010/products/1

## TODO
- More tests
- Handling nil pointers
- handle date from transport

## Tutorials
- Resource practice tutorial for using go, TDD and postgres:

    https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

