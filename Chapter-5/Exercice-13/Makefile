MODULE_NAME=GoExercices/Chapter-5/Exercice-13

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/findlinks

clean:
	rm -f ${GOPATH}/bin/findlinks
	#chmod -R u+w pkg ; rm -rf pkg

