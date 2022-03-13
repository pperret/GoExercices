MODULE_NAME=GoExercices/Chapter-9/Exercice-6

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/mandelbrot

test:
	go mod tidy
	GOMAXPROCS=1 go test -bench=. $(MODULE_NAME)/mandelbrot -benchtime=30s
	GOMAXPROCS=2 go test -bench=. $(MODULE_NAME)/mandelbrot -benchtime=30s
	GOMAXPROCS=4 go test -bench=. $(MODULE_NAME)/mandelbrot -benchtime=30s
	GOMAXPROCS=8 go test -bench=. $(MODULE_NAME)/mandelbrot -benchtime=30s
	GOMAXPROCS=16 go test -bench=. $(MODULE_NAME)/mandelbrot -benchtime=30s


clean:
	rm -f ${GOPATH}/bin/mandelbrot
