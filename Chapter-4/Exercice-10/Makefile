MODULE_NAME=dummy.com/C4E10

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/issues

go.mod:
	go mod init $(MODULE_NAME)

run:
	${GOPATH}/bin/issues repo:golang/go is:open json decoder

clean:
	rm -f ${GOPATH}/bin/issues
	rm -f go.mod
