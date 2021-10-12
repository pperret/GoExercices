MODULE_NAME=GoExercices/Chapter-7/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/treesort

clean:
	rm -f ${GOPATH}/bin/treesort
