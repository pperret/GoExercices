@echo off
set GOPATH=%cd%
go test -bench=. popcount_test
set GOPATH=

