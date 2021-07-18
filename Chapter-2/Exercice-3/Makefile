MODULE_NAME=dummy.com/C2E3

test: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go test -bench=. $(MODULE_NAME)/popcount

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -f src/go.mod
