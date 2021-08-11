MODULE_NAME=Chapter-4/Exercice-10

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
