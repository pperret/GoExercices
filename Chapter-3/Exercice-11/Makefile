MODULE_NAME=GoExercices/Chapter-3/Exercice-11

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/comma

clean:
	rm -f ${GOPATH}/bin/comma
