package main

/* Write a program to sort an array of integers. The program should partition the array into 4 parts, 
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
 Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array
 should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/
   
import (
	"fmt"
	"strings"
	"strconv"
)


func main() {
	var strFromUser string

	fmt.Printf("enter sequence of up to 12 integers (seperated by spaces): ")
	n, err := fmt.Scanf("%q", &strFromUser)
	fmt.Printf("Got %d characters\n", n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Here are the user's numbers: " + strFromUser)

	sliceOfStrs := strings.Split(strFromUser, " ")
	n = len(sliceOfStrs)
	fmt.Println(n)
	fullSliceOfInts := make([]int,0,n)	
	for _, s := range sliceOfStrs {
		newInt, _ := strconv.Atoi(s)
		fullSliceOfInts = append(fullSliceOfInts, newInt)
	}
	// slices dicing tricks from https://tour.golang.org/concurrency/2
	sliceofInts1 := fullSliceOfInts[:n/4] // make([]int,0,n/4)
	sliceofInts2 := fullSliceOfInts[n/4:n/2] // make([]int,0,n/4)
	sliceofInts3 := fullSliceOfInts[n/2:n*3/4] // make([]int,0,n/4)
	sliceofInts4 := fullSliceOfInts[n*3/4:] // make([]int,0,n/4)

	// slices dicing tricks from https://tour.golang.org/concurrency/2
	fmt.Println("s1", sliceofInts1)
	fmt.Println("s2", sliceofInts2)
	fmt.Println("s3", sliceofInts3)
	fmt.Println("s4", sliceofInts4)
	// fmt.Println("Test Swap",sliceofInts1)
	//Swap(sliceofInts, 3)

	fmt.Println("now sort the whole thing", sliceofInts1)
	c := make(chan []int)
	go BubbleSort(sliceofInts1, c)
	go BubbleSort(sliceofInts2, c)
	go BubbleSort(sliceofInts3, c)
	go BubbleSort(sliceofInts4, c)

	sliceofInts1Sorted := make([]int,0,n/4)
	sliceofInts1Sorted = <- c
	fmt.Println("result of slice1", sliceofInts1Sorted)
	
	sliceofInts2Sorted := make([]int,0,n/4)
	sliceofInts2Sorted = <- c
	fmt.Println("result of slice2", sliceofInts2Sorted)
	
	sliceofInts3Sorted := make([]int,0,n/4)
	sliceofInts3Sorted = <- c
	fmt.Println("result of slice3", sliceofInts3Sorted)

	sliceofInts4Sorted := make([]int,0,n/4)
	sliceofInts4Sorted = <- c
	fmt.Println("result of slice4", sliceofInts4Sorted)
}

func BubbleSort(sliceToSort []int, c chan []int){
	fmt.Println(sliceToSort)
	for i, _:= range sliceToSort {
		if !sweep(sliceToSort, i) {
			// fmt.Println(sliceToSort)
			c <- sliceToSort
		}
	}
}
func sweep(sliceToSort []int, prevPasses int) bool {
	didSwap := false
	for i := 0; i < len(sliceToSort) - prevPasses; i++ {
		if len(sliceToSort) > i+1 {
			// fmt.Println("comparing: ", sliceToSort[i], sliceToSort[i+1])
			if sliceToSort[i] > sliceToSort[i+1] {
				Swap(sliceToSort, i)
				didSwap = true
			}
		}
	}
	return didSwap
}
func Swap(sints []int, indx int) {
	// swap the int at position indx with the one at indx+1
	intToSwap := sints[indx]
	sints[indx] = sints[indx+1]
	sints[indx+1] = intToSwap
	// fmt.Println("Swapped: ", sints[indx], intToSwap)
}