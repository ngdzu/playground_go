package mypackage

import (
	"flag"
	"fmt"
)

func usingflag() {

	// if specify -o, the bool flag is turned on
	// or specify -o=false to explicitly set the value of the flag
	minusO := flag.Bool("o", false, "usage of o desc")
	minusC := flag.Bool("c", false, "usage of c description")

	// if use -k, must supply the value, or not using -k at all
	// to get the default value 0
	minusK := flag.Int("k", 0, "usage of k description")

	flag.Parse()

	fmt.Println("-o:", *minusO)
	fmt.Println("-c:", *minusC)
	fmt.Println("-k:", *minusK)

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
