MODULE_NAME=GoExercices/Chapter-8/Exercice-1

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/clock
	go install $(MODULE_NAME)/clockwall

clean:
	rm -f ${GOPATH}/bin/clock
	rm -f ${GOPATH}/bin/clockwall
