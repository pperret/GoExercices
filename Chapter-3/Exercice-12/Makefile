MODULE_NAME=GoExercices/Chapter-3/Exercice-12

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/anagram

clean:
	rm -f ${GOPATH}/bin/anagram
