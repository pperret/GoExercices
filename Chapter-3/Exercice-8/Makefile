MODULE_NAME=dummy.com/C3E8

test: src/go.mod
	cd src ; go mod tidy
	cd src ; GOPATH="${PWD}" go test -bench=. $(MODULE_NAME)/mandelbrot

src/go.mod:
	cd src ; go mod init $(MODULE_NAME)

clean:
	rm -f src/go.mod
