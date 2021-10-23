MODULE_NAME=GoExercices/Chapter-7/Exercice-4

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/stringreader

clean:
	rm -f ${GOPATH}/bin/stringreader
