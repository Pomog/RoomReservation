package main

import "testing"

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("run() failed")
	}
}

/*
go test -coverprofile=coverage.out && go tool cover -html coverage.out
*/
