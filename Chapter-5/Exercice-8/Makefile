MODULE_NAME=GoExercices/Chapter-5/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/findattribute

clean:
	rm -f ${GOPATH}/bin/findattribute
	#chmod -R u+w pkg ; rm -rf pkg
