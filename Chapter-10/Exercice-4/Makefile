MODULE_NAME=GoExercices/Chapter-10/Exercice-4

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/depends

	go mod download golang.org/x/sys
	go mod download golang.org/x/net
	go mod download golang.org/x/term
	go mod download golang.org/x/text

clean:
	rm -f ${GOPATH}/bin/depends
