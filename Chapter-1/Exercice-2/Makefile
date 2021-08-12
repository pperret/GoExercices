MODULE_NAME=GoExercices/Chapter-1/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/echo

clean:
	rm -f ${GOPATH}/bin/echo
