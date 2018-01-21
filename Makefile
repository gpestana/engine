.PHONY: build test sample

build:
	go build 

test:
	go get .
	go test -cover ./...

sample:
	./engine -config=./sample/conf.json
