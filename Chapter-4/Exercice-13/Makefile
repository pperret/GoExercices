MODULE_NAME=dummy.com/C4E13

build: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go install $(MODULE_NAME)/omdb

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -rf bin
	rm -f src/go.mod