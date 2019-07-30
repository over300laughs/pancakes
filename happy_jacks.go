package main

import "fmt"

const ServeThemHappy = "+"
const ServeThemSad = "-"

//main is simply an entry point and driver
func main(){

	// INPUT: 
	// I really don't like having 5 as an input so I'm ommitting this. 
	// We already know input length dynamically from the slice length so let's not add more to maintain. :-)

	//creating static inputs
	input := make([]string,0)
	allFlips := make([]int,0)

	// Inputs can simply be added here and the program will format it gracefully
	input = append(input, "-", "-+", "+-", "+++", "--+-")

	// Printing Static Inputs
	printInput(input)
	
	//Find the number of flips for each input
	for _, stack := range input {
		flips := getFlips(stack)
		allFlips = append(allFlips, flips)
	}

	//printing output via static function
	printOutput(allFlips)

	return

}

// getFlips is a small unexported function that simply examines the state of each pancake
// It then counts each time a variable (pancakeSide) changes state. 
// If the state changes then we just add one flip.
// Its important to note that the lambda value of the variable pancakeSide
// is what determines which side of the pancake we're counting
// If, for example the Amercian Flag was flown at half mast for a day,
// and the boss wanted to only serve sad sided pancakes,
// we could simply change the lambda state to be the const ServeThemSad 
// and we could count flips to make all the pancakes sad for the day.
func getFlips(flapJacks string) int {
	flips := 0

	// We want our customers to have happy pancakes!
	// Change this to ServeThemSad (AKA "-") to count the flips to serve sad pancakes :-(
	pancakeSide := ServeThemHappy

	// Simple loop for checking number of flips
	// X-ray pancake vision to see the state of each pancake
	for i := len(flapJacks); i > 0; i-- {
		
		// If this expression is true we simply need to count the state changes of the pancakeSide variable
		if flapJacks[i-1:i] != pancakeSide {

			// Save what we're trying find out!
			flips++

			// Set the value so we only add a flip when the state of pancakeSide changes
			pancakeSide = flapJacks[i-1:i] 
		}
	}

	return flips
} 

// printInput is unexported and used simply to see the input for the challenge.
// Error handling has been ommitted due to simplicity of the needs of this challenge.
// In production errors would have been handled to reflect the same behavior as the rest of the repository
// Its also important to note that I'm aware that the challenge does not specifiy to see the input but
// I wanted to add it for slightly more accessability to the end user
func printInput(input []string) {
	for i := 0; i < len(input); i++ {
		fmt.Printf("Case #%d Input: %s \n", i+1, input[i])
	}
}

// printOutput is unexported and used simply to see the output for the challenge.
// Error handling has been ommitted due to simplicity of the needs of this challenge.
// In production errors would have been handled to reflect the same behavior as the rest of the repository
func printOutput(flipSlice []int) {
	for key, flips := range flipSlice {
		key++ //0 based so we need to increase it by one
		fmt.Printf("Case #%d: %d \n", key, flips)
	}
}
