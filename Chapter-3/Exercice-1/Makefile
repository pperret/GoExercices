MODULE_NAME=GoExercices/Chapter-3/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/surface

clean:
	rm -f ${GOPATH}/bin/surface
