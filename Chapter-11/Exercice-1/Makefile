MODULE_NAME=GoExercices/Chapter-11/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/charcount

test:
	go mod tidy
	go test $(MODULE_NAME)/charcount

clean:
	rm -f ${GOPATH}/bin/charcount
