package main

import ("testing")

/*
Purpose: To show that Golang is capable of producing 
higher order functions and the typechecker can make sure
they are being used properly.
*/
	
// Test case that adds two values from the higher order functions

	func TestSum(t *testing.T) {

		x := DoubleAndHalf(5)   //returns 12
		y := HalfAndMultBy3(9)	// uses 4 returns 6	

		total := x() + y()

		if total != 18 {

		   t.Errorf("Value is incorrect, got: %d, expected: %d.", total, 18)
		   
		}
	}


//
	func TestName(t *testing.T) {

		tempName := "My name is First Last"
		
		fullName := func(f_name string, l_name string) string{

			return "My name is " + f_name + " " + l_name
		}
		
		check := firstLastName(fullName)
	
		if check != tempName {
			t.Errorf("Strings do not match.)
		}
	
	}