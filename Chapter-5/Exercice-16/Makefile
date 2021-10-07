MODULE_NAME=GoExercices/Chapter-5/Exercice-16

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/join

clean:
	rm -f ${GOPATH}/bin/join
