MODULE_NAME=dummy.com/C2E4

ifndef GOPATH
$(error GOPATH is not set)
endif

test: go.mod
	go mod tidy
	go test -bench=. $(MODULE_NAME)/popcount

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f go.mod
