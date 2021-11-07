MODULE_NAME=GoExercices/Chapter-7/Exercice-11

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/database

clean:
	rm -f ${GOPATH}/bin/database
