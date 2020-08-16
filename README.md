# go-web-benchmark

- [echo](https://echo.labstack.com/)
- [fiber](https://gofiber.io/)
- [gin](https://gin-gonic.com/)

## Make

```
make
```

## Run

```
bin/echo
bin/fiber
bin/gin
```

## Load Test

```
make bench

# or
make wrk

# or
make hey
```

use [hey](https://github.com/rakyll/hey)

```
go get -u github.com/rakyll/hey
```

run

```
hey -n 100000 http://127.0.0.1:8080/ping
```
or
```
hey -z 5s http://127.0.0.1:8080/ping
```

## Tools

use [curlie](https://curlie.io/)

```
go get -u -v github.com/rs/curlie
```

```
curlie localhost:8080/ping
```

```
HTTP/1.1 200 OK
Date: Sun, 16 Aug 2020 19:24:11 GMT
Content-Type: application/json
Content-Length: 18

{
    "message": "pong"
}
```
