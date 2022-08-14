MODULE_NAME=GoExercices/Chapter-11/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:

test:
	go mod tidy
	go test $(MODULE_NAME)/split

clean:
	rm -f ${GOPATH}/bin/split
