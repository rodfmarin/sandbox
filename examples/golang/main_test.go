package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("this is a test, testing")
}

func TestGetRandNames(t *testing.T) {
	res := GetRandNames(0)
	if len(res) != 0 {
		t.Error("Expected no names, got ", len(res))
	}

	res = GetRandNamesImpl(1, []string{"solid"}, []string{"snake"})

	if len(res) != 1 {
		t.Error("Expected a name array containing 'solid snake', got %v", res[0])
	}

	if res[0] != "solid snake" {
		t.Error("Expected a name array containing 'solid snake', got %v", res[0])
	}
}

func TestGetRandName(t *testing.T) {
  res := GetRandName([]string{"solid"}, []string{"snake"})
	if res != "solid snake" {
		t.Error("Expected a name array containing 'solid snake', got %v", res)
	}
}

func TestGetRandNamePure(t *testing.T) {
	res := GetRandNamePure([]string{"A","B"}, []string{"1","2"},
		func(n int) int{
			return 0
		})
		if res != "A 1" {
			t.Error("Expected a name array containing 'A 1', got %v", res)
		}
	res = GetRandNamePure([]string{"A","B"}, []string{"1","2"},
		func(n int) int{
			return 1
		})
		if res != "B 2" {
			t.Error("Expected a name array containing 'B 2', got %v", res)
		}

}
