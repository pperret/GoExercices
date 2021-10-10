MODULE_NAME=GoExercices/Chapter-6/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/intset

clean:
	rm -f ${GOPATH}/bin/intset
