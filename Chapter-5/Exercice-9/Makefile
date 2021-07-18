MODULE_NAME=dummy.com/C5E9

test: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go test -v $(MODULE_NAME)/expand

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -rf bin
	rm -f src/go.mod

