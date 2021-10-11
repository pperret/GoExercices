MODULE_NAME=GoExercices/Chapter-7/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/counter

clean:
	rm -f ${GOPATH}/bin/counter
