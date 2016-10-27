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

  res = GetRandNamesImpl(1,[]string{"solid"},[]string{"snake"})
  if len(res) != 1  && res[0] != "solid snake"  {
    t.Error("Expected a name array containing 'solid snake', got %v", res[0])
  }

}

