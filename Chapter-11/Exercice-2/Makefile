MODULE_NAME=GoExercices/Chapter-11/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/intset

test:
	go mod tidy
	go test $(MODULE_NAME)/intset

clean:
	rm -f ${GOPATH}/bin/intset
