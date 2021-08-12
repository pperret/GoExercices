MODULE_NAME=GoExercices/Chapter-4/Exercice-10

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/issues

run:
	${GOPATH}/bin/issues repo:golang/go is:open json decoder

clean:
	rm -f ${GOPATH}/bin/issues
