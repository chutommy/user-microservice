# Account (Software Design Document)
### Tommy Chu
### August 16, 2020

## Introduction
The Account is a microservice, which provides all necessary actions to manipulate account records in the database. The client should be able to call the service with JSON formatted messages through HTTP or HTTPS. Each account record holds the ID, username, hashed password, email address, phone number, first name, last name, birth day, permanent address, mailing address, the timestamp of the account's creation and update. The service should be completely independent and isolated.

### Goals
    - basic CRUD operations
    - login (email/password or username/password validation)
    - read by ID or parameters
    - soft/hard delete
    - authentication for all commands

### Non-Goals
    - assigning roles
    - determine relationships
    - separating rights and privileges

## Development
The Account service will be mainly written in Go and SQL - database queries. The API endpoints will be managed and served with the Go's Gin framework. The whole service and database is going to be containerized with the Docker and for the API documentation will be used swag-go tool.

All the configuration can be modified in config.yml file (database connection credentials, server settings - timeouts, exposed ports). Each endpoint will be independently tested with the unit tests.

## API Endpoints
    - `POST /auth` Authenticate the account's email and the hashed password (in request's body).

    - `POST /accounts/` Create a new account, data - inside request's body.
    - `GET /accounts/` Get all accounts.
    - `GET /accounts/{id}` Get an account by ID.
    - `GET /accounts/params` Get an account by paramaters (query parameters, e.g. /q?fname=Tommy&lname=chu&bday=16042003&email=tommychu2256@gmail.com).
    - `UPDATE /accounts/{id}` Update an account.
    - `DELETE /accounts/{id}` Delete an account (soft delete).

## Milestones
Event | Completed
----- | ---------
Database |
Data service |
Configuration |
Controller |
Server |
Documentation |
Containerization |
Finish |
