MODULE_NAME=dummy.com/C2E2

build: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go install $(MODULE_NAME)/testconv

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -rf bin
	rm -f src/go.mod
 