MODULE_NAME=dummy.com/C3E5

ifndef GOPATH
$(error GOPATH is not set)
endif

build: go.mod
	go mod tidy
	go install $(MODULE_NAME)/mandelbrot

go.mod:
	go mod init $(MODULE_NAME)

clean:
	rm -f ${GOPATH}/bin/mandelbrot
	rm -f go.mod