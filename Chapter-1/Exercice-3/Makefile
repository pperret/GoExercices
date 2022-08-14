MODULE_NAME=GoExercices/Chapter-1/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:
	go mod tidy
	go install $(MODULE_NAME)/echo
	
test:
	go mod tidy
	go test -bench=. $(MODULE_NAME)/echo -args tutu tata titi

clean:
	# Done
