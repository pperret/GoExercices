MODULE_NAME=dummy.com/C4E10

build: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go install $(MODULE_NAME)/issues

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

run:
	bin/issues repo:golang/go is:open json decoder

clean:
	rm -rf bin
	rm -f src/go.mod
