package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	expected := []string{"1", "2", "Fizz", "4", "Buzz"}
	res := fizzBuzz(5)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
