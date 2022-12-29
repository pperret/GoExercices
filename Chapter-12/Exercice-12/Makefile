MODULE_NAME=GoExercices/Chapter-12/Exercice-12

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/server

test:
	#go mod tidy
	#go test -v $(MODULE_NAME)/params

clean:
	rm -f ${GOPATH}/bin/server
