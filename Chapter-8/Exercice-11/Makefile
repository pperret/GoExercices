MODULE_NAME=GoExercices/Chapter-8/Exercice-11

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/fetch

clean:
	rm -f ${GOPATH}/bin/fetch
