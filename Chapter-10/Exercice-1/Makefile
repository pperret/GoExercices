MODULE_NAME=GoExercices/Chapter-10/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/imgconv

clean:
	rm -f ${GOPATH}/bin/imgconv
