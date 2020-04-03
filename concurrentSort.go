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
	"sort"
	"sync"
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

	// fmt.Println("s1", sliceofInts1)
	// fmt.Println("s2", sliceofInts2)
	// fmt.Println("s3", sliceofInts3)
	// fmt.Println("s4", sliceofInts4)

	fmt.Println("now sort the whole thing", sliceofInts1, sliceofInts2, sliceofInts3, sliceofInts4)
	// since sorting works in-place, use the Wait approach for Sync, versus the Channel approach
	var wg sync.WaitGroup
	wg.Add(4)
	go stdSort(sliceofInts1, &wg)
	go stdSort(sliceofInts2, &wg)
	go stdSort(sliceofInts3, &wg)
	go stdSort(sliceofInts4, &wg)
	wg.Wait()
	var fullArray []int
	fullArray = append(fullArray,sliceofInts1... )
	fullArray = append(fullArray,sliceofInts2... )
	fullArray = append(fullArray,sliceofInts3... )
	fullArray = append(fullArray,sliceofInts4... )
	fmt.Println("full partially sorted: ", fullArray)
	sort.Ints(fullArray)
	fmt.Println("fully sorted: ", fullArray)
}

func stdSort(sliceToSort []int, wg *sync.WaitGroup) []int{
	defer wg.Done()
	fmt.Println("in conc stdSort: ", sliceToSort)
	sort.Ints(sliceToSort)
	return sliceToSort
} 