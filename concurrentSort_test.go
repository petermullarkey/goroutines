package main

/* Test the program by running it twice and testing it with a different sequence of integers each time. 
The first test sequence of integers should be all positive numbers 
and the second test should have at least one negative number.
*/

import ("testing"
		"reflect"
)

func Test(t *testing.T) {
     var tests = []struct {
     	 s []int; want []int
	}{
		{[]int{6, 3, 7, 4, 9, 10, 2, 1, 4, 7, 43, 5}, []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 9, 10, 43}},
		{[]int{9, 2, 6, 1, -2, 8, 7, 4, 9, 10, 2, 1}, []int{-2, 1, 1, 2, 2, 4, 6, 7, 8, 9, 9, 10}},
	}
	for _, c := range tests {
		retVal := sortConcurrentController(c.s)
	    if !reflect.DeepEqual(retVal, c.want) {
	       t.Errorf("sortConcurrentController(%q) wanted %q got %v", c.s, c.want, retVal)
	    }
	}
}