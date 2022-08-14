MODULE_NAME=GoExercices/Chapter-11/Exercice-7

ifndef GOPATH
$(error GOPATH is not set)
endif

build:

test:
	go mod tidy
	go test -bench=. -benchmem -timeout=0 $(MODULE_NAME)/intset -args -max=1000 -count=10
	go test -bench=. -benchmem -timeout=0 $(MODULE_NAME)/intset -args -max=1000000 -count=10
	go test -bench=. -benchmem -timeout=0 $(MODULE_NAME)/intset -args -max=1000000 -count=1000
	go test -bench=. -benchmem -timeout=0 $(MODULE_NAME)/intset -args -max=1000000000 -count=1000
	go test -bench=. -benchmem -timeout=0 $(MODULE_NAME)/intset -args -max=1000000000 -count=1000000

clean:
	rm -f ${GOPATH}/bin/intset
