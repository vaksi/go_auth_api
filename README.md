# go_auth_api
Golang Restfull Api with gin-gonic framework and jwt(json web token)

## GO Simple Rest API Authentication with JWT (json web token)
 This repository is example basic auth with JWT for golang http://gopkg.in/appleboy/gin-jwt.v2 and golang framework gin-gonic https://godoc.org/github.com/gin-gonic/gin

## Running and Testing
Don't forget for setup your GOPATH and GOROOT. After that you must be import golang package.
```sh
$ go get gopkg.in/appleboy/gin-jwt.v2
$ go get github.com/fvbock/endless
$ go get github.com/gin-gonic/gin
```

### Build and running
you must be build and running your apps
```sh
$ go build
$ go run
```
### Testing 
Registering for new user
```sh
$ curl -v --form username=admin --form password=admin http://localhost:3100/register
```
```sh
*   Trying ::1...
* Connected to localhost (::1) port 3100 (#0)
> POST /register HTTP/1.1
> Host: localhost:3100
> User-Agent: curl/7.43.0
> Accept: */*
> Content-Length: 250
> Expect: 100-continue
> Content-Type: multipart/form-data; boundary=------------------------1167045fb0275d4e
> 
< HTTP/1.1 100 Continue
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 29 Dec 2016 23:31:55 GMT
< Content-Length: 70
< 
{"status":"posted","users":[{"username":"admin","password":"admin"}]}
* Connection #0 to host localhost left intact
```
Login user
```sh
$ curl -H "Content-Type: application/json" -X POST -d '{"username":"admin","password":"admin"}' http://localhost:3100/login
```
```sh
{"expire":"2016-12-30T07:39:43+07:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODMwNTgzODMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTQ4MzA1NDc4M30.-szphYikAlrvWj150xpH_Vqce8w1ijr80DVL45vTXRo"}
```
Testing using auth token
```sh
$ curl -H "Authorization: Bearer XXXyourtokenXXX" -X GET http://localhost:3100/auth/hello
```
```sh
{"text":"Hello World."}
```
