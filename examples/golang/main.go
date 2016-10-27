package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello Worl")
}

var firstNames = []string{"first"}
var lastNames = []string{"last"}

func GetRandNames(numNames int) []string{
  return GetRandNamesImpl(numNames,firstNames, lastNames)
}

func GetRandNamesImpl(numNames int, firstNames []string, lastNames []string) []string{
  return make([]string,numNames)
}
