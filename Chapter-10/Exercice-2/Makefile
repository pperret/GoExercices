MODULE_NAME=GoExercices/Chapter-10/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/arlist

clean:
	rm -f ${GOPATH}/bin/arlist
