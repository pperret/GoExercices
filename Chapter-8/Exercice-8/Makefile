MODULE_NAME=GoExercices/Chapter-8/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/reverb

clean:
	rm -f ${GOPATH}/bin/reverb
