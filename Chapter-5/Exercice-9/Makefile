MODULE_NAME=Chapter-5/Exercice-9

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
