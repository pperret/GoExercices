MODULE_NAME=GoExercices/Chapter-9/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

test:
	go mod tidy
	go test $(MODULE_NAME)/bank

clean:
	#rm -f ${GOPATH}/bin/bank
