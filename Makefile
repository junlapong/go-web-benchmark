build:
	go build -o bin/echo echo/main.go
	go build -o bin/fiber fiber/main.go
	go build -o bin/gin gin/main.go

bench:
	wrk -c 100 -t 1 -d 5s http://localhost:8080/

test:
	http :8080/

clean:
	rm -rf bin/

