MODULE_NAME=GoExercices/Chapter-9/Exercice-4

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/pipeline

clean:
	rm -f ${GOPATH}/bin/pipeline
