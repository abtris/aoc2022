package main

import "testing"

func TestGetHungryElfCalories(t *testing.T) {
	data := prepareData("input_test")
	expected := 24000
	real := getHungryElfCalories(data)
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}

}

func TestGetHungryLastThreeElfCalories(t *testing.T) {
	data := prepareData("input_test")
	expected := 45000
	real := getHungryLastThreeElfCalories(data)
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}

}
