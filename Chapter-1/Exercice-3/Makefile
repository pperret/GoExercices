MODULE_NAME=Chapter-1/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

test: go.mod
	go mod tidy
	go test -bench=. $(MODULE_NAME)/echo -args tutu tata titi

go.mod:
	go mod init $(MODULE_NAME)

clean:
	# Done
