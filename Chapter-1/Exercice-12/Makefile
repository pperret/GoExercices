MODULE_NAME=GoExercices/Chapter-1/Exercice-12

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/server

clean:
	rm -f ${GOPATH}/bin/server
 