package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Create a piece of data of a random size
	piece := createpiece(9999)
	fmt.Println("\n Unsorted \n", piece)
	// Create var wg to create a wait group for goroutine communication
	var wg sync.WaitGroup
	// Retrieve time before the quicksort executes
	now := time.Now()
	// Initiate quicksort algorithm
	quicksort(piece)
	// Retrieve time after the quick sort algorithm
	post := time.Now()
	fmt.Println("\n The time before quicksort: ", now)
	fmt.Println("\n The time after quicksort: ", post)
	// Display sorted data
	fmt.Println("\n Sorted \n", piece)
	// Create a new piece of random data
	newPiece := createpiece(9999)
	// Display the unsorted data
	fmt.Println("\n Unsorted \n", newPiece)
	wg.Add(1)
	// Grab the time before the bubblesort
	preBubblesort := time.Now()
	// bubblesort
	go bubblesort(newPiece, &wg)
	// Wait for goroutines to complete
	wg.Wait()
	// Grab the time after the bubblesort
	postBubblesort := time.Now()
	//  Display info
	fmt.Println("\n Time before bubblesort: ", preBubblesort)
	fmt.Print("\n Time post bubblesort: ", postBubblesort)
	fmt.Println("\n Sorted \n", newPiece)
	diffQuicksort := post.Sub(now)
	diifBubblesort := postBubblesort.Sub(preBubblesort)
	// Displays the time difference between the algorithms in nanoseconds
	fmt.Println("\n Time Taken for Quicksort: ", diffQuicksort.Nanoseconds())
	fmt.Println("\n Time Taken for Bubblesort: ", diifBubblesort.Nanoseconds())
}

// Create a piece of data func
func createpiece(size int) []int {
	// Create a variable piece of an array of ints
	piece := make([]int, size, size)
	// Fill the array with random numbers
	for i := 0; i < size; i++ {
		// It's important to note that for the method of random generation I used, that taking 2 random numbers between
		// 0 - 999 ensures that the range between each number in the list remains between -999 - 999 to ensure symmetry
		piece[i] = rand.Intn(999) - rand.Intn(999)
	}
	return piece
}

// Quicksorting Function
func quicksort(a []int) []int {
	// If the length of the data piece is less than 2 i.e the Index 0, then there is no need
	// to sort and the code should just return that element unchanged
	if len(a) < 2 {
		return a
	}

	// Create two vars, from the first index 0 to the last index len(a) - 1 (to account for arrays starting at 0)
	left, right := 0, len(a)-1

	// This will locate a number that is roughly in the centre of the array
	// We do not necessarily need to find the real centre although it would save a marginal
	// ammount of time when dealing with small data sets
	centre := rand.Int() % len(a)

	// Here we swap the values for the right and centre indices
	// to ensure that our arrays 'pivot' is at one end, making the sort more efficient
	a[centre], a[right] = a[right], a[centre]

	// For loop to iterate through the array, disregarding the value at the point of index i each time
	for i, _ := range a {
		// If the index is less than the pivot, it's value should be moved left
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			// We increment the left index to ensure that we aren't overwriting the value we just stored
			left++
		}
	}
	// Swap values of the pivot point and the last value reached before the If statement
	// fails
	a[left], a[right] = a[right], a[left]

	// Recursively apply quicksort to a slice up to the new pivot point
	quicksort(a[:left])
	// Recursively apply quicksort to a slice starting from the next index after the pivot point
	quicksort(a[left+1:])

	return a
}

// Bubblesort Function
func bubblesort(a []int, wg *sync.WaitGroup) []int {
	// Validation check to see if array is long enough to sort
	// If not, return the only element in the array
	if len(a) < 2 {
		return a
	}

	// Decrement counter when goroutine finishes
	defer wg.Done()

	// Loop to iterate through the array
	for i := 0; i < len(a)-1; i++ {
		// Loop for comparing each element in the array to it's neighbour
		for j := 0; j < len(a)-i-1; j++ {
			// Check if the current element is less than the next
			if a[j] < a[j+1] {
				// If a is less than the next index in the array, swap their positions
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a

}
