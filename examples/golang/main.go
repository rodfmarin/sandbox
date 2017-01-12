package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(GetRandNames(1)[0])
}

// var firstNames = []string{"first", "aasdfasd", "232432"}
// var lastNames = []string{"last", "testteststests", "thisisalastname"}
 var firstNames = FileToLines("firstnames.txt")
 var lastNames = FileToLines("lastnames.txt")


func GetRandNames(numNames int) []string {
	return GetRandNamesImpl(numNames, firstNames, lastNames)
}

func GetRandNamesImpl(numNames int, firstNames []string, lastNames []string) []string {
	/*This function takes a number of times to get a random name (first and last) and returns
	  an array of string
	  Make a call to GetRandName n times
	*/
	names := make([]string, numNames)
	for ii, _ := range names {
		tempName := GetRandName(firstNames, lastNames)
		names[ii] = tempName
	}
	return names
}

func GetRandName(firstNames []string, lastNames []string) string {

	//return firstNames[rand.Intn(len(firstNames))] + " " + lastNames[rand.Intn(len(lastNames))]
	return GetRandNamePure(firstNames, lastNames, rand.Intn)
}

type RandInt func(int) int

func GetRandNamePure(firstNames []string, lastNames []string, r RandInt) string {

	return firstNames[r(len(firstNames))] + " " + lastNames[r(len(lastNames))]
}

func FileToLines(filepath string) []string {
	f,err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	//does defer do automatic releasing of this resource?
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return lines

}
