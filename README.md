[![GoDoc](https://godoc.org/github.com/bruno-chavez/go-web-template?status.svg)](https://godoc.org/github.com/bruno-chavez/go-web-template)
[![Build Status](https://travis-ci.org/bruno-chavez/go-web-template.svg?branch=master)](https://travis-ci.org/bruno-chavez/go-web-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruno-chavez/go-web-template)](https://goreportcard.com/report/github.com/bruno-chavez/go-web-template)

# Description

`go-web-template` main purpose is to be a starting point 
for web development in Go, 
usually just as a API for a front-end application.

# How to use

1. Rename package, project, go mod file and delete .git directory

2. Download dependencies

3. Have Redis installed and a server running.

4. Have PostgresSQL installed and a database created.

5. Rename the `.env.example` file to `.env`  and customize the parameters accordingly

Note: Don't store the `SESSION_STORE_KEY` in your source code and ensure your key is sufficiently random and large.

# Features 

+ Ready to use custom authentication routes, for registering, 
login in and login out users.

+ Ready to use cookie based sessions with 
[gorilla/sessions](https://github.com/gorilla/sessions) and 
Redis as a session store with 
[redistore](https://github.com/boj/redistore).

+ Ready to use Relational Database connection to store user 
and domain data with PostgreSQL,
[lib/pq](https://github.com/lib/pq) as the driver and 
[sqlx](https://github.com/jmoiron/sqlx) 
to help with raw queries.

+ Ready to handle CORS requests with 
[rs/cors](https://github.com/rs/cors).

+ Safely hashes and salts user passwords with the official 
[bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) 
implementation.

+ Fast and easy to use router with 
[julienschmidt/httprouter](https://github.com/julienschmidt/httprouter).

+ Migration file for a `user` table. 
Currently no automatic way of running migrations is provided.

+ Loads environment variables from an `.env` file 
with the help of [godotenv](https://github.com/joho/godotenv).

+ Dependency management with Go Modules.

+ Continuous Integration with Travis-CI.

# To Do

+ Middleware

+ Unit testing.

# Contribute

Found a bug or an error? Post it in the 
[issue tracker](https://github.com/bruno-chavez/go-web-template/issues).

Want to add an awesome new feature? 
[Fork](https://github.com/bruno-chavez/go-web-template/fork) 
this repository, add your feature on a new branch, 
then send a pull request.

# License
The MIT License (MIT)
Copyright (c) 2019 Bruno Chavez