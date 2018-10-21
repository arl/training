package main

import "testing"

func TestMyFunction(t *testing.T) {
	hello := "hello"
	if "hello" != hello {
		t.Fail()
	}
}
