MODULE_NAME=GoExercices/Chapter-7/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/countingwriter

clean:
	rm -f ${GOPATH}/bin/countingwriter
