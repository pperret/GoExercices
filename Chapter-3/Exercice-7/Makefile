MODULE_NAME=GoExercices/Chapter-3/Exercice-7

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/newton

clean:
	rm -f ${GOPATH}/bin/newton
