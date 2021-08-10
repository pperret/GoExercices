MODULE_NAME=dummy.com/C5E2

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/countelementtype

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/countelementtype
	#chmod -R u+w pkg ; rm -rf pkg
	rm -f go.mod go.sum
