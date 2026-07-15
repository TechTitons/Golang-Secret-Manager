run:
	go run ./cmd/server

build:
	go build -o secret-manager ./cmd/server

test:
	go test ./...

docker-build:
	docker build -t golang-secret-manager .

docker-run:
	docker run -p 8080:8080 golang-secret-manager

clean:
	rm -f secret-manager