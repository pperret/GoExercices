MODULE_NAME=GoExercices/Chapter-5/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/printtextnodes

clean:
	rm -f ${GOPATH}/bin/printtextnodes
	#chmod -R u+w pkg ; rm -rf pkg
