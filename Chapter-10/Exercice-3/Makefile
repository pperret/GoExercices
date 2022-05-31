MODULE_NAME=GoExercices/Chapter-10/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/fetch

run: build
	${GOPATH}/bin/fetch http://gopl.io/ch1/helloworld?go-get=1 | grep go-import

clean:
	rm -f ${GOPATH}/bin/fetch
