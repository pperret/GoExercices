MODULE_NAME=GoExercices/Chapter-8/Exercice-7

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/extractor

clean:
	rm -f ${GOPATH}/bin/extractor
