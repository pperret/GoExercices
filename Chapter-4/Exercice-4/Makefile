MODULE_NAME=GoExercices/Chapter-4/Exercice-4

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/rotate

clean:
	rm -f ${GOPATH}/bin/rotate
