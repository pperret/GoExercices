MODULE_NAME=GoExercices/Chapter-5/Exercice-2

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/countelementtype

clean:
	rm -f ${GOPATH}/bin/countelementtype
	#chmod -R u+w pkg ; rm -rf pkg
