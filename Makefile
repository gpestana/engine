.PHONY: build test sample
ifndef VERBOSE
.SILENT:
endif


build:
	go build 

test:
	go get .
	go test -cover ./...

sample:
	./engine -config=./sample/conf.json
