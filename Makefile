GOPATH=$(CURDIR)/godeps/

default: godeps build

godeps: 
	env GOPATH="${GOPATH}" go get

build:
	env GOPATH="${GOPATH}" go build -o bin/go-aws-mon

deps:
	
clean:
		rm -f bin/*
