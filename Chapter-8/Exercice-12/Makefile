MODULE_NAME=GoExercices/Chapter-8/Exercice-12

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/chat

clean:
	rm -f ${GOPATH}/bin/chat
