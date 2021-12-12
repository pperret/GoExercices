MODULE_NAME=GoExercices/Chapter-8/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/ftpd

clean:
	rm -f ${GOPATH}/bin/ftpd
