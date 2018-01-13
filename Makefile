build:
	go build 

test:
	go get .
	go test -cover ./...
