MODULE_NAME=GoExercices/Chapter-9/Exercice-3

ifndef GOPATH
$(error GOPATH is not set)
endif

build:

test:
	go mod tidy
	# MallocNanoZone is required on Mac to prevent crashes with race option (Monterey, go 1.17.5)
	MallocNanoZone=0 go test -run=TestConcurrent -race -v $(MODULE_NAME)/memo

clean:
	
