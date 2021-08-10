MODULE_NAME=dummy.com/C4E4

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/rotate

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/rotate
	rm -f go.mod