MODULE_NAME=GoExercices/Chapter-12/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/display

test:
	go mod tidy
	go test -v $(MODULE_NAME)/display

clean:
	rm -f ${GOPATH}/bin/display
