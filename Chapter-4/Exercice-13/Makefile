MODULE_NAME=GoExercices/Chapter-4/Exercice-13

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/omdb

clean:
	rm -f ${GOPATH}/bin/omdb
