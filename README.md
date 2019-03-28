[![GoDoc](https://godoc.org/github.com/bruno-chavez/go-web-template?status.svg)](https://godoc.org/github.com/bruno-chavez/go-web-template)
[![Build Status](https://travis-ci.org/bruno-chavez/go-web-template.svg?branch=master)](https://travis-ci.org/bruno-chavez/go-web-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruno-chavez/go-web-template)](https://goreportcard.com/report/github.com/bruno-chavez/go-web-template)

# Description

`go-web-template` main purpose is to be a starting point 
for a back-end web service developed in Go, usually just as a 
REST API for any front-end project in Vue/React/Angular.

# Features 

+ Session management with 
[gorilla/sessions](https://github.com/gorilla/sessions) and 
Redis as session store with 
[redistore](https://github.com/boj/redistore).

+ Ready to use Relational Database to store user and domain data with PostgreSQL,
[lib/pq](https://github.com/lib/pq) as the driver and 
[sqlx](https://github.com/jmoiron/sqlx) to help with the raw queries.

+ Safely hashes and salts user passwords with [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt), no more

+ Fast and easy to use router with [httprouter](https://github.com/julienschmidt/httprouter).

+ Ready to handle CORS requests with [rs/cors](https://github.com/rs/cors).

+ Loads environment variables from an `.env` file with the help of [godotenv](https://github.com/joho/godotenv).

+ Dependancy management with Go Modules.

+ Continuous Integration with Travis-CI.

+ Each route has access to the sessions store 
and the db connection pool thanks to dependency injection 
(simply reference `c.Db` or `c.SessionStore` once 
your handlers are methods of the `Controller` struct).

# To Do

+ Implement HTTPS with [certmagic](https://github.com/mholt/certmagic).

+ Database migrations with either 
[sql-migrate](https://github.com/rubenv/sql-migrate) or 
[migrate](https://github.com/golang-migrate/migrate).

+ Unit testing.

# Contribute

Found an bug or an error? Post it in the [issue tracker](https://github.com/bruno-chavez/go-web-template/issues).

Want to add an awesome new feature? [Fork](https://github.com/bruno-chavez/go-web-template/fork) this repository, add your feature on a new branch, then send a pull request.

# License
The MIT License (MIT)
Copyright (c) 2019 Bruno Chavez