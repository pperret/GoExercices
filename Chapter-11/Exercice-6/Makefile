MODULE_NAME=GoExercices/Chapter-11/Exercice-6

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/popcount

test:
	go mod tidy
	go test -bench=. -benchmem -benchtime=5s $(MODULE_NAME)/popcount -args -value=0 -expected=0
	go test -bench=. -benchmem -benchtime=5s $(MODULE_NAME)/popcount -args -value=1 -expected=1
	go test -bench=. -benchmem -benchtime=5s $(MODULE_NAME)/popcount -args -value=3 -expected=2
	go test -bench=. -benchmem -benchtime=5s $(MODULE_NAME)/popcount -args -value=7 -expected=3
	go test -bench=. -benchmem -benchtime=5s $(MODULE_NAME)/popcount -args -value=15 -expected=4

clean:
	rm -f ${GOPATH}/bin/popcount
