MODULE_NAME=GoExercices/Chapter-4/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/digest

clean:
	rm -f ${GOPATH}/bin/digest
