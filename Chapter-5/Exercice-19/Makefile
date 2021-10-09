MODULE_NAME=GoExercices/Chapter-5/Exercice-19

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/panic

clean:
	rm -f ${GOPATH}/bin/panic
