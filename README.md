# Kode Test Task

Anagram storage written in golang with postgresql.

Basically, this is a simple anagram storage, that enables logged in users to store and query anagrams.
This storage uses a feature of anagrams, that if two words are anagrams, those words would be equal if they are sorted.
e.g. ["dog", "god"] after sorting would result in ["dgo", "dgo"]

Anagrams are stored in two one-to-many related tables.
One stores sorted versions of anagrams.
Other stores all anagrams and their relation to a sorted one.
On query a JOIN is performed.

### Prerequisites

```
docker
docker-compose
```

### Installation

Just clone it and run docker-compose up

```
git clone https://github.com/kembo91/kode-test-task
cd kode-test-task
docker-compose up
```

Web interface runs at your localhost:8080

You can play with this storage via a web interface, or example curl commands are given below.

## Running the tests

Tests are written for a database with the use of [go-txdb](https://github.com/DATA-DOG/go-txdb/)
They require a running postgresql instance with set up user (postgres), password (postgres) and database (test)

```
cd kode-test-task/test
go test ./...
```

## API

## Create new user

### Request

`POST /api/signup`

User signup creates a new user in a database and responds with a cookie with jwt token on success

```
curl -i -H 'Content-Type: application/json' -X POST -d '{"Username":"newusername","Password":"newpassword"}' http://localhost:8080/api/signup
```

### Response

```
HTTP/1.1 202 Accepted
Content-Type: application/json
Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk3MDc1fQ.Ei_lgvd89a7LKje9R6TZIRAZm6ig5d80hTFa1FuV9Go; Path=/api; Expires=Sun, 10 Jan 2021 16:44:35 GMT
Date: Sun, 10 Jan 2021 16:34:35 GMT
Content-Length: 21

{"result":"success"}
```

## Sign in existing user

### Request

`POST /api/signin`

User signin looks for provided user credentials in a database and returns a jwt token on success

```
curl -i -H 'Content-Type: application/json' -X POST -d '{"Username":"newusername","Password":"newpassword"}' http://localhost:8080/api/signin
```

### Response

```
HTTP/1.1 202 Accepted
Content-Type: application/json
Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk3MzU1fQ.LY2l5HzWBSJk6USDbl0cBPRHawjQmOIJdbr6qqTmAEg; Path=/api; Expires=Sun, 10 Jan 2021 16:49:15 GMT
Date: Sun, 10 Jan 2021 16:39:15 GMT
Content-Length: 21

{"result":"success"}
```

## Sign out

### Request

`GET /api/signout`

User signout overrides client cookie with an expired one

```
curl -i -H "Content-Type: application/json" -X GET http://localhost:8080/api/signout
```

### Response

```
HTTP/1.1 202 Accepted
Content-Type: application/json
Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InJhbmRvbSIsImV4cCI6MTYxMDI5NjM5Mn0.WuOLLOmoR1AVC6U3LOcB_Y6PDpMuORsRzU15BqBo90g; Path=/api; Expires=Sun, 10 Jan 2021 16:33:12 GMT
Date: Sun, 10 Jan 2021 16:43:12 GMT
Content-Length: 21

{"result":"success"}
```

## Insert anagram

### Request

`POST /api/anagram/insert`

Inserts anagram into the database.
Avaliable only for signed in users with a valid token cookie
Also sets a new cookie with a prolonged expiration time

```
curl -i -H 'Content-Type: application/json' -X POST -d '{"Query" : "dog"}' -b 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk3ODkxfQ.4oglJDQ4t-V76Ag6u77-5XckBnkt4VKx0VX1GC58Ies' http://localhost:8080/api/anagram/insert
```

### Response

```
HTTP/1.1 202 Accepted
Content-Type: application/json
Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk4MDAyfQ.9YguDcZetY0_bVk1lePSzb9gYsfW_FwG-5XpgyqKySI; Path=/api; Expires=Sun, 10 Jan 2021 16:33:12 GMT
Date: Sun, 10 Jan 2021 16:43:12 GMT
Content-Length: 21

{"result":"success"}
```

## Retrieve anagram

### Request

`POST /api/anagram/retrieve`

Returns all anagrams of a query anagram, found in a database.
Avaliable only for signed in users with a valid token cookie.
Also sets a new cookie with a prolonged expiration time.

```
curl -i -H 'Content-Type: application/json' -X POST -d '{"Query" : "dog"}' -b 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk3ODkxfQ.4oglJDQ4t-V76Ag6u77-5XckBnkt4VKx0VX1GC58Ies' http://localhost:8080/api/anagram/retrieve
```

### Response

```
Content-Type: application/json
Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk4MzYwfQ.kWXg83L_pTHl460EVrZQP2CNwUViV3x1qzgLIvc8Lc0; Path=/api; Expires=Sun, 10 Jan 2021 17:06:00 GMT
Date: Sun, 10 Jan 2021 16:56:00 GMT
Content-Length: 20

["god","dog","odg"]
```

## Retrieve all anagrams

### Request

`GET /api/anagram/retrieve`

Returns all anagrams, found in a database.
Avaliable only for signed in users with a valid token cookie.
Also sets a new cookie with a prolonged expiration time.

```
curl -i -H "Content-Type: application/json" -b 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk3ODkxfQ.4oglJDQ4t-V76Ag6u77-5XckBnkt4VKx0VX1GC58Ies' -X GET http://localhost:8080/api/retrieve
```

### Response

```
Content-Type: application/json
Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ld3VzZXJuYW1lIiwiZXhwIjoxNjEwMjk4NjY3fQ.F9xk-uuyQ3dvacJ2FH9-xrIbo65QPw1nub3blF2RTQA; Path=/api; Expires=Sun, 10 Jan 2021 17:11:07 GMT
Date: Sun, 10 Jan 2021 17:01:07 GMT
Content-Length: 20

["dog","odg","god"]
```

## Time consumption

Dev environment setup ~3h

Database code ~4h

Router handlers code ~4h

Test environment setup ~2h

After a test env setup I had to refactor a lot ~4h

Writing tests ~4h

Web interface ~3h

Deployment ~6h
