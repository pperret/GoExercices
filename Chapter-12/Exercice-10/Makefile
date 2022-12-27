MODULE_NAME=GoExercices/Chapter-12/Exercice-10

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/sexpr

test:
	go mod tidy
	go test -v $(MODULE_NAME)/sexpr

clean:
	rm -f ${GOPATH}/bin/sexpr
