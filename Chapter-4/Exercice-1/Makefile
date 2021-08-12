MODULE_NAME=GoExercices/Chapter-4/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build: 
	go mod tidy
	go install $(MODULE_NAME)/sha256

clean:
	rm -f ${GOPATH}/bin/sha256
