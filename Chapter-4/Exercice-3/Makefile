MODULE_NAME=GoExercices/Chapter-4/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/reverse

clean:
	rm -f ${GOPATH}/bin/reverse
