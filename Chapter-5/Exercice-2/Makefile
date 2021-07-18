MODULE_NAME=dummy.com/C5E2

build: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go install $(MODULE_NAME)/countelementtype

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -rf bin
	chmod -R u+w pkg ; rm -rf pkg
	rm -f src/go.mod src/go.sum
