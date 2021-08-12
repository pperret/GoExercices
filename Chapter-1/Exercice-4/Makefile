MODULE_NAME=GoExercices/Chapter-1/Exercice-4

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/dup

clean:
	rm -f ${GOPATH}/bin/dup
