package main

import (
	
	
    "time"
	"math/rand"
	"sort"
)

//sorts the numbers in an array 
func performSort(size int){

	//Seed to generate random numbers
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

 //makes list of a certain size 
	list := make([]int, size)


	//Adds random values between 1 and 5000 into array
	for i := 0; i < len(list); i++ {  
            list[i] = r1.Intn(5000)
		}

       //Prints the unsorted array
       //fmt.Println("Unsorted",list)

		//sorts the array
		sort.Ints(list[:])	

	   //Prints the sorted array
       //fmt.Println("Sorted",list)

			
}







