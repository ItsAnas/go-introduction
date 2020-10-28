package main

import "testing"

func TestFiboBasic(t *testing.T) {
	res := fibo(1)
	expected := 1
	if res != expected {
		t.Error("Expected ", expected, " Got ", res)
	}
}

func TestFiboAdvanced(t *testing.T) {
	res := fibo(6)
	expected := 8
	if res != expected {
		t.Error("Expected ", expected, " Got ", res)
	}
}

func TestFiboError(t *testing.T) {
	res := fibo(-1)
	expected := 0
	if res != expected {
		t.Error("Expected ", expected, " Got ", res)
	}
}
