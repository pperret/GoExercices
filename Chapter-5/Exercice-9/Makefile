MODULE_NAME=dummy.com/C5E9

ifndef GOPATH
$(error GOPATH is not set)
endif

test: go.mod
	go mod tidy
	go test -v $(MODULE_NAME)/expand

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/expand
	rm -f go.mod