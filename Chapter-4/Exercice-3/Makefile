MODULE_NAME=dummy.com/C4E3

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/reverse

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/reverse
	rm -f go.mod