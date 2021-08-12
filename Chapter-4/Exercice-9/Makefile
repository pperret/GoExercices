MODULE_NAME=GoExercices/Chapter-4/Exercice-9

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/wordfreq

clean:
	rm -f ${GOPATH}/bin/wordfreq
