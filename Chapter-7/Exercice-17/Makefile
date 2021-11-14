MODULE_NAME=GoExercices/Chapter-7/Exercice-17

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/xmlselect

clean:
	rm -f ${GOPATH}/bin/xmlselect
