MODULE_NAME=GoExercices/Chapter-13/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/bzipper

test:
	go mod tidy
	go test -v $(MODULE_NAME)/bzip

clean:
	rm -f ${GOPATH}/bin/bzipper
