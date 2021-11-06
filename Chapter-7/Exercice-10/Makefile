MODULE_NAME=GoExercices/Chapter-7/Exercice-10

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/palindrome

clean:
	rm -f ${GOPATH}/bin/palindrome
