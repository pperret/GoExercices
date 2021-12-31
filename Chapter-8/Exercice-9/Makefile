MODULE_NAME=GoExercices/Chapter-8/Exercice-9

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/du

clean:
	rm -f ${GOPATH}/bin/du
