MODULE_NAME=GoExercices/Chapter-5/Exercice-15

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/minmax

clean:
	rm -f ${GOPATH}/bin/minmax
