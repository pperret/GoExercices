MODULE_NAME=GoExercices/Chapter-13/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/equal

test:
	go mod tidy
	go test -v $(MODULE_NAME)/equal

clean:
	rm -f ${GOPATH}/bin/equal
