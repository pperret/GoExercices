MODULE_NAME=dummy.com/C1E2

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/echo

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/echo
	rm -f go.mod
