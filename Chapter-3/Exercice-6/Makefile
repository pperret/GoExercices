MODULE_NAME=GoExercices/Chapter-3/Exercice-6

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/mandelbrot

clean:
	rm -f ${GOPATH}/bin/mandelbrot
