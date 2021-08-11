MODULE_NAME=Chapter-5/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/findattribute

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/findattribute
	#chmod -R u+w pkg ; rm -rf pkg
