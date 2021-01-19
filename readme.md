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

## TODO
 - use transactions in database queries
 - use context through all request flow


## Goose usage
add migration example:
```
 goose -dir ./migrations create initial sql
 goose -dir ./migrations postgres "user=tasksuser password=password123431 dbname=tasks sslmode=disable" up
```

## Check it out using Postman
here is a link to postman collection:
https://www.getpostman.com/collections/88178d2e6247dadef6b4


## Athorization
To make request authorized it should have `Authorization` header with `base64` encoder username there. Usename can contain only letters and numbers there.