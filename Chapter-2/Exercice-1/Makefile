MODULE_NAME=GoExercices/Chapter-2/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/testconv

clean:
	rm -f ${GOPATH}/bin/testconv
 