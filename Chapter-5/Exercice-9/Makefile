MODULE_NAME=GoExercices/Chapter-5/Exercice-9

ifndef GOPATH
$(error GOPATH is not set)
endif

test:
	go mod tidy
	go test -v $(MODULE_NAME)/expand

clean:
	rm -f ${GOPATH}/bin/expand
