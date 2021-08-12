MODULE_NAME=GoExercices/Chapter-5/Exercice-12

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/outline

clean:
	rm -f ${GOPATH}/bin/outline
	#chmod -R u+w pkg ; rm -rf pkg

