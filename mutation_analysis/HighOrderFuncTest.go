package main

//import ("testing")


// manbearpig cannot run mutation tests on Go test files so the only way to get it to perform
// an analysis was to comment out the test features of the file such as import "testing" and 	
// t *testing.T used in the functions. 


	func TestSum( /* t *testing.T*/ ) {

		x := DoubleAndHalf(5)   //returns 12
		y := HalfAndMultBy3(4)	//returns 6	

		total := x() + y()

		if total != 18 {

		   t.Errorf("Value is incorrect, got: %d, expected: %d.", total, 18)
		   
		}
	}



	func TestName( /* t *testing.T */ ) {

		tempName := "My name is First Last"
		
		fullName := func(f_name string, l_name string) string{

			return "My name is " + f_name + " " + l_name
		}
		
		check := firstLastName(fullName)
	
		if check != tempName {
			t.Errorf("Strings do not match.")
		}
	
	}