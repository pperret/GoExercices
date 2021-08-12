MODULE_NAME=GoExercices/Chapter-4/Exercice-14

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/issueshtml

clean:
	rm -f ${GOPATH}/bin/issueshtml
