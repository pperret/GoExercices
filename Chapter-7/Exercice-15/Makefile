MODULE_NAME=GoExercices/Chapter-7/Exercice-15

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/calc

clean:
	rm -f ${GOPATH}/bin/calc
