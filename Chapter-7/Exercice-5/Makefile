MODULE_NAME=GoExercices/Chapter-7/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/limitreader

clean:
	rm -f ${GOPATH}/bin/limoitreader
