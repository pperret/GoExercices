MODULE_NAME=GoExercices/Chapter-4/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/unique

clean:
	rm -f ${GOPATH}/bin/unique
