package main

import "testing"

func TestAdd(t *tesing.T) {
	if Add(5,2) != 7 {
		t.Error("5+2 != 7")
	}
} 