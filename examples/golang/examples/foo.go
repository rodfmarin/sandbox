package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	names := make([]string, 10)
	rand.Seed(time.Now().UTC().UnixNano())
	for ii, _ := range names {
		fmt.Printf("ii=%v\n", ii)
	}

	fnames := []string{
		"rod",
		"kyle",
		"alan",
		"colin",
	}
	lnames := []string{
		"marin",
		"burton",
		"kwan",
		"fortuner",
		"cumberbatch",
	}
	fmt.Printf("pickRandElementForReal(%v) => %v\n",
		fnames, pickRandElementForReal(fnames))
	var res string
	res = GetRandName(fnames, lnames, pickFirstElement)
	fmt.Printf("in the constant case: %v\n", res)
	res = GetRandName([]string{"solid"}, []string{"snake"}, pickRandElementForReal)
	fmt.Printf("in the rand case (of 1): %v\n", res)
	res = GetRandName(fnames, lnames, pickRandElementForReal)
	fmt.Printf("in the full monty case: %v\n", res)

	// assert "rod marin"   == GetRandName(fnames, lnames, pickFirstElement)
	// assert "kyle burton" == GetRandName(fnames, lnames, pickSecondElement)
}

type RandFnType func([]string) string

func GetRandName(fnames []string, lnames []string, randFn RandFnType) string {
	return randFn(fnames) + " " + randFn(lnames)
}

func pickFirstElement(things []string) string {
	return things[0]
}

func pickSecondElement(things []string) string {
	return things[1]
}

func pickRandElementForReal(things []string) string {
	return things[rand.Intn(len(things))]
}
