package main


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

