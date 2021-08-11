MODULE_NAME=Chapter-3/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

test: go.mod
	go mod tidy
	go test -bench=. $(MODULE_NAME)/mandelbrot

go.mod:
	go mod init $(MODULE_NAME)

clean:
	# Done
