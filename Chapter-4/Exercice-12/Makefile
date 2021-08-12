MODULE_NAME=GoExercices/Chapter-4/Exercice-12

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/xkcd

clean:
	rm -f ${GOPATH}/bin/xkcd
