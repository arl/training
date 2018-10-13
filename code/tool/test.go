package main

import "testing"

func TestMyFunction(t *testing.T) {
	if "hello" == "world" {
		t.Fail()
	}
}
