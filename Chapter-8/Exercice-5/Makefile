MODULE_NAME=GoExercices/Chapter-8/Exercice-5

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/mandelbrot

test:
	go mod tidy
	go test -bench=. $(MODULE_NAME)/mandelbrot -benchtime=30s


clean:
	rm -f ${GOPATH}/bin/mandelbrot
