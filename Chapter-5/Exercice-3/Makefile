MODULE_NAME=dummy.com/C5E3

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/printtextnodes

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/printtextnodes
	#chmod -R u+w pkg ; rm -rf pkg
	rm -f go.mod go.sum
