# Account (Software Design Document)

## Tommy Chu

### August 16, 2020

## Introduction

The Account is a microservice, which provides all necessary actions
to manipulate account records in the database. The client should
be able to call the service with JSON formatted messages through
HTTP or HTTPS. Each account record holds the ID, username,
hashed password, email address, phone number, first name, last name,
birth day, permanent address, mailing address, the timestamp of the
account's creation and update. The service should be completely
independent and isolated.

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

The Account service will be mainly written in Go
and SQL - database queries. The API endpoints will be
managed and served with the Go's Gin framework. The
whole service and database is going to be containerized
with the Docker and for the API documentation will
be used swag-go tool.

All the configuration can be modified in config.yml
file (database connection credentials,
server settings - timeouts, exposed ports). Each
endpoint will be independently tested with the unit tests.

All collumns are optional, except for the ID, Email,
CreatedAt and DeletedAt. ID, Username, Email and Phone
must be unique. The ID should have the UUID data type.

The Postgres database will be run in the Docker container
and the data should persist inside the db/data directory.
When the container is initialized for the first time,
the init.sql should run to create a table.

## API Endpoints

    - `POST /auth` Authenticate the account's email and
    the hashed password (in request's body).
    - `POST /accounts/` Create a new account, data - inside request's body.
    - `GET /accounts/` Get all accounts.
    - `GET /accounts/{id}` Get an account by ID.
    - `GET /accounts/params` Get an account by paramaters (query parameters,
    e.g. /q?fname=Tommy&lname=chu&bday=16042003&email=tommychu2256@gmail.com).
    - `UPDATE /accounts/{id}` Update an account.
    - `DELETE /accounts/{id}` Delete an account (soft delete).

## Steps

Event | Completed
----- | ---------
Database | August 17, 2020
Data service | August 20, 2020
Controller | August 23, 2020
Configuration | August 25, 2020
Server | August 31, 2020
Documentation | September 4, 2020
Containerization | September 7, 2020
Finish |
