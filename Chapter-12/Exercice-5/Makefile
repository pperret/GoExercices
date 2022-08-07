MODULE_NAME=GoExercices/Chapter-12/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/encoder

test:
	go mod tidy
	go test -v $(MODULE_NAME)/encoder

clean:
	rm -f ${GOPATH}/bin/encoder
