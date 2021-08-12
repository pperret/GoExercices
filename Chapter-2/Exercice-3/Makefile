MODULE_NAME=GoExercices/Chapter-2/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

test:
	go mod tidy
	go test -bench=. $(MODULE_NAME)/popcount

clean:
	# Done
