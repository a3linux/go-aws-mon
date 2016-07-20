GOPATH=$(CURDIR)/godeps/

default: godeps build

godeps: 
	env GOPATH="${GOPATH}" go get

build:
	env GOOS=linux GOPATH="${GOPATH}" go build -o bin/go-aws-mon

deps:
	
clean:
		rm -f bin/*
