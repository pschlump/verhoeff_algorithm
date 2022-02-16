//
// Copyright (C) Philip Schlump, 2014-2016
// MIT Licensed.
//
package verhoeff_algorithm

import (
	"fmt"
	"testing"
)

func Test_VerhoeffAlgorithm(t *testing.T) {
	tests := []struct {
		Str      string
		Verhoeff int
	}{
		{Str: "75872", Verhoeff: 2},                  // From PHP example
		{Str: "12345", Verhoeff: 1},                  //
		{Str: "142857", Verhoeff: 0},                 //
		{Str: "123456789012", Verhoeff: 0},           //
		{Str: "8473643095483728456789", Verhoeff: 2}, //
		{Str: "501", Verhoeff: 1},                    // Generated by C code in ./test
		{Str: "8102439", Verhoeff: 6},                //
		{Str: "0112113", Verhoeff: 1},                //
		{Str: "2314231", Verhoeff: 6},                //
		{Str: "4423123", Verhoeff: 9},                //
		{Str: "55123123", Verhoeff: 7},               //
		{Str: "3332433", Verhoeff: 0},                //
		{Str: "883478", Verhoeff: 9},                 //
		{Str: "884378347", Verhoeff: 7},              //
		{Str: "748384783783738101", Verhoeff: 6},     //
		{Str: "58584", Verhoeff: 9},                  //
	}

	for ii, test := range tests {
		r := GenerateVerhoeff(test.Str)
		if r != test.Verhoeff {
			t.Errorf("Test %d: Expected %d got %d\n", ii, test.Verhoeff, r)
		}

		// test the string
		rs := GenerateVerhoeffString(test.Str)
		exp := fmt.Sprintf("%s%d", test.Str, r)
		if rs != exp {
			t.Errorf("Test %d: Expected %s got %s\n", ii, exp, rs)
		}

		// test the validate -- should indicate true
		b := ValidateVerhoeff(rs)
		if !b {
			t.Errorf("Test %d: Expected 'true' got %v\n", ii, b)
		}

		// change a char -- restest validate - should indicate false
		rs = GenerateVerhoeffString(test.Str)
		// ms := rs[1:2] + rs[0:1] + rs[2:] // create an error
		ms := rs[1:2] + "6" + rs[2:] // create an error
		b = ValidateVerhoeff(ms)
		if b {
			t.Errorf("Test %d: Expected 'false' got %v\n", ii, b)
		}

	}
}
