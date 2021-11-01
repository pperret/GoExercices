MODULE_NAME=GoExercices/Chapter-7/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/multisort

clean:
	rm -f ${GOPATH}/bin/multisort
