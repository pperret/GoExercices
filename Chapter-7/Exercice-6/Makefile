MODULE_NAME=GoExercices/Chapter-7/Exercice-6

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/tempflag

clean:
	rm -f ${GOPATH}/bin/tempflag
