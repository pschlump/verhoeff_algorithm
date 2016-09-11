package main

//
// Verhoeff Algorithm
//
// Copyright (C) Philip Schlump, 2014-2016
// MIT Licensed.
//
// Command line to run the Verhoeff_Algorithm
//
// To generate checksums
//
// 		$ ./cli 123 124
//
// Output:
//
// 		123 => 1233
// 		124 => 1246
//
// To validate
//
// 		$ ./cli -v 1233 1236
//
// Output:
//
//		1233 valid
//		1236 INVALID!
//
import (
	"flag"
	"fmt"

	"github.com/pschlump/verhoeff_algorithm"
)

var Verify = flag.Bool("verify", false, "Verify each string") // 1
func init() {
	flag.BoolVar(Verify, "v", false, "Verify each string") // 1
}

func main() {
	flag.Parse()
	fns := flag.Args()

	for _, s := range fns {
		if *Verify {
			if verhoeff_algorithm.ValidateVerhoeff(s) {
				fmt.Printf("%s valid\n", s)
			} else {
				fmt.Printf("%s INVALID!\n", s)
			}
		} else {
			t := verhoeff_algorithm.GenerateVerhoeffString(s)
			fmt.Printf("%s => %s\n", s, t)
		}
	}
}
