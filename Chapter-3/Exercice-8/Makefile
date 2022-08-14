MODULE_NAME=GoExercices/Chapter-3/Exercice-8

ifndef GOPATH
$(error GOPATH is not set)
endif

build:

test:
	go mod tidy
	go test -bench=. $(MODULE_NAME)/mandelbrot

clean:
	# Done
