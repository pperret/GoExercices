set GOPATH=%cd%
go test -bench=. echo -args tutu tata titi
set GOPATH=

