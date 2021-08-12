MODULE_NAME=GoExercices/Chapter-1/Exercice-6

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/lissajous

clean:
	rm -f ${GOPATH}/bin/lissajous
