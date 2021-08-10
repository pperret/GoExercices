MODULE_NAME=dummy.com/C3E11

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/comma

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/comma
	rm -f go.mod