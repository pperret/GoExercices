MODULE_NAME=GoExercices/Chapter-5/Exercice-7

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/prettyprint

test:
	go mod tidy
	go test -v $(MODULE_NAME)/prettyprint

clean:
	rm -f ${GOPATH}/bin/prettyprint
	#chmod -R u+w pkg ; rm -rf pkg

