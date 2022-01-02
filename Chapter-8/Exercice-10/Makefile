MODULE_NAME=GoExercices/Chapter-8/Exercice-10

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/crawl

clean:
	rm -f ${GOPATH}/bin/crawl
