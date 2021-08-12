MODULE_NAME=GoExercices/Chapter-4/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/charcount

clean:
	rm -f ${GOPATH}/bin/charcount
