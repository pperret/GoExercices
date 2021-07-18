MODULE_NAME=dummy.com/C4E14

build: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go install $(MODULE_NAME)/issueshtml

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -rf bin
	rm -f src/go.mod