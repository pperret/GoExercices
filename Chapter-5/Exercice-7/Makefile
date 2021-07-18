MODULE_NAME=dummy.com/C5E7

build: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go install $(MODULE_NAME)/prettyprint

test: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go test -v $(MODULE_NAME)/prettyprint

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -rf bin
	chmod -R u+w pkg ; rm -rf pkg
	rm -f src/go.mod src/go.sum

