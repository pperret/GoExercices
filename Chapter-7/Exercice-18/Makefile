MODULE_NAME=GoExercices/Chapter-7/Exercice-18

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/xmltree

clean:
	rm -f ${GOPATH}/bin/xmltree
