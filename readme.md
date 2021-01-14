# Description
This is a project for yalantis golang school.  
Task description is here: https://docs.google.com/document/d/1PPAbDVllQYpw7bFRStGB_Gbcoj2TRs9NvyPfepuAf8w/edit#

## TODO 
 - use goose for migrations
 - use request validators in request handlers
 - use context through all request flow


## Goose usage
add migration example:
```
 goose -dir ./migrations create initial sql
```

## Athorization
To make request authorized it should have `Authorization` header with `base64` encoder username there. Usename can contain only letters and numbers there.