MODULE_NAME=GoExercices/Chapter-5/Exercice-17

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/findtag

clean:
	rm -f ${GOPATH}/bin/findtag
