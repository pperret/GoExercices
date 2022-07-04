MODULE_NAME=GoExercices/Chapter-11/Exercice-4

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/palindrome

test:
	go mod tidy
	go test $(MODULE_NAME)/palindrome

clean:
	rm -f ${GOPATH}/bin/palindrome
