# League Challenge

## Description

This is an project which implements end-points for league coding challenge. Aplication that contains enpoints according to requirements backend chanllege.
It exposes endpoints to post a matrix by a file.

## Overview
This application was designed following clean architecture and uses the principles of Robert Martin. It contains endpoints for:
- Echo a matrix
- Flatten a matrix
- Invert a matrix
- Multiply a matrix
- Sum a matrix


Rule of Clean Architecture by Uncle Bob

- Independent of Frameworks.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html


This project has 4 Domain layer :

- Entities Layer
- Usecases Layer
- Adapters Layer
- Frameworks Layer

## Commands
- Runt Tests: `make test`
- Run Application: `make run`
- Build Application: `make build`

## Send request
```
curl -F 'file=@/path/file.csv' "localhost:8080/api/v1/echo"
curl -F 'file=@/path/file.csv' "localhost:8080/api/v1/flatten"
curl -F 'file=@/path/file.csv' "localhost:8080/api/v1/invert"
curl -F 'file=@/path/file.csv' "localhost:8080/api/v1/multiply"
curl -F 'file=@/path/file.csv' "localhost:8080/api/v1/sum"
```