package main

import ("testing")



func TestFib(t *testing.T) {


	num := 10
    fibNum := 55

if num <= 0 {    // Prints error message if negative number is used since fib cant work on negative values

	t.Errorf("Negatve number was used. Try value greater than 0")


} else {

      trial := fib(num)

  if trial != 55{  //passes with 55, fails with any other number

		t.Errorf("Fibonachi of %d is %d", num , fibNum)
	}

 }

}

/*
Output when it passes

$ go test
PASS
ok  


Output when it fails

$ go test
--- FAIL: TestFib (0.00s)
    REPL_test.go:22: Fibonachi of 10 is 55
FAIL
exit status 1
FAIL 


*/


	   
	
