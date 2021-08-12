MODULE_NAME=GoExercices/Chapter-5/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/countwords

clean:
	rm -f ${GOPATH}/bin/countwords
	#chmod -R u+w pkg ; rm -rf pkg
