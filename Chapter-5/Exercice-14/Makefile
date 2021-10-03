MODULE_NAME=GoExercices/Chapter-5/Exercice-14

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/breadthfirst

clean:
	rm -f ${GOPATH}/bin/breadthfirst
