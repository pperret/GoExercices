
build:
	GOPATH="${PWD}" go get golang.org/x/net/html
	GOPATH="${PWD}" go install prettyprint

test:
	GOPATH="${PWD}" go get golang.org/x/net/html
	GOPATH="${PWD}" go test -v prettyprint

clean:
	rm -rf bin
	rm -rf pkg
	rm -rf src/golang.org
