MODULE_NAME=GoExercices/Chapter-5/Exercice-10

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/toposort

clean:
	rm -f ${GOPATH}/bin/toposort


