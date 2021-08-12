MODULE_NAME=GoExercices/Chapter-1/Exercice-10

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/fetchall

clean:
	rm -f ${GOPATH}/bin/fetchall
