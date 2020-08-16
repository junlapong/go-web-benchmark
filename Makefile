build:
	go build -o bin/echo echo/main.go
	go build -o bin/fiber fiber/main.go
	go build -o bin/gin gin/main.go

bench:
	ab -k -c 10 -n 100000 http://localhost:8080/ping

wrk:
	wrk -c 10 -t 1 -d 5s http://localhost:8080/ping

hey:
	hey -c 10 -cpus 1 -z 5s http://localhost:8080/ping

test:
	http :8080/ping

clean:
	rm -rf bin/

