MODULE_NAME=GoExercices/Chapter-8/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/netcat

clean:
	rm -f ${GOPATH}/bin/netcat
