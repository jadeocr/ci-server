dev:
	go build -o ci-server.sh cmd/github.com/jadeocr/ci-server/main.go
	./ci-server.sh

build:
	go build -o ci-server.sh cmd/github.com/jadeocr/ci-server/main.go
