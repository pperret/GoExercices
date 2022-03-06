MODULE_NAME=GoExercices/Chapter-9/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/pingpong

clean:
	rm -f ${GOPATH}/bin/pingpong
