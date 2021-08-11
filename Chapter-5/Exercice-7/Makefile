MODULE_NAME=Chapter-5/Exercice-7

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/prettyprint

test: go.mod
	go mod tidy
	go test -v $(MODULE_NAME)/prettyprint

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/prettyprint
	#chmod -R u+w pkg ; rm -rf pkg

