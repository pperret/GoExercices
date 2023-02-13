MODULE_NAME=GoExercices/Chapter-13/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/cycle

test:
	go mod tidy
	go test -v $(MODULE_NAME)/cycle

clean:
	rm -f ${GOPATH}/bin/cycle
