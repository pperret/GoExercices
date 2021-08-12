MODULE_NAME=GoExercices/Chapter-4/Exercice-11

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/issues

clean:
	rm -f ${GOPATH}/bin/issues
