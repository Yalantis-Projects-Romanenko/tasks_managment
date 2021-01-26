# Description
This is a project for yalantis golang school.  
Task description is here: https://docs.google.com/document/d/1PPAbDVllQYpw7bFRStGB_Gbcoj2TRs9NvyPfepuAf8w/edit#

## Start the project

open root directory of the project in command line and run:

```
docker-compose up .
cd api
go get -u github.com/pressly/goose/cmd/goose
goose -dir ./migrations postgres "user=tasksuser password=password123431 dbname=tasks sslmode=disable" up
go run main.go
```

## Tests

run

```
docker-compose -f docker-compose.test.yml up -d
go test ./api
```

## Goose usage

add migration example:

```
 goose -dir ./migrations create initial sql
 goose -dir ./migrations postgres "user=tasksuser password=password123431 dbname=tasks sslmode=disable" up
 goose -dir ./migrations postgres "user=tasksuser password=password123431 dbname=tasks sslmode=disable" down-to 20210114124747
```

## Linter Usage

```
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2
golangci-lint --version
cd api
golangci-lint run
```

cd

## Check it out using Postman
here is a link to postman collection:
https://www.getpostman.com/collections/88178d2e6247dadef6b4

## Authorization
To make request authorized, it should have `Authorization` header with `base64` encoded username there. Username can contain only letters and numbers.


## TODO
- handle not found requests
- cover with tests
