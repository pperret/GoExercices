MODULE_NAME=GoExercices/Chapter-10/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/jpeg

clean:
	rm -f ${GOPATH}/bin/jpeg
