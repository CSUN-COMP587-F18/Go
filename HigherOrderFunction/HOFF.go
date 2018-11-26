package main

import("fmt")


 /*A higher order function that  multilplies a number by itself 
 and then returns a function that returns the number divided by 2)
 */

func DoubleAndHalf(x int) func() int {
	
	x = x * x
	
		return func() int {
	  
	  x = x / 2

	  return x
	   }
	}
	

/*Another higher order function that halves a number and then
 returns a function that returns the number multiplied by 3
 */

  func HalfAndMultBy3(x int) func() int { 

	x = x / 2 
	  
		return func() int {
	  
		x = x * 3 
		
	    return x
	  }
  }

 //Another way to perform a higher order function 
  func firstLastName(f func(string , string) string) string{

	name := f("First","Last")
	
	return name
	
	}

	func main(){

		fullName := func(f_name string, l_name string) string{

			return "My name is " + f_name + " " + l_name
		}
		
		fullString := firstLastName(fullName)

		fmt.Println(fullString)
	}