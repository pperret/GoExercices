MODULE_NAME=Chapter-2/Exercide-4

ifndef GOPATH
$(error GOPATH is not set)
endif

test: go.mod
	go mod tidy
	go test -bench=. $(MODULE_NAME)/popcount

go.mod:
	go mod init $(MODULE_NAME)

clean:
	# Done
