package errors_test

import (
	"errors"
	"testing"
	"fmt"
	
)

 
//Tests to see if allocated data in a variable is the same if data in variable
//are stored in another variable
func TestEquality(t *testing.T) {
	
	text1:= errors.New("Hello")
	text2 := text1

	if text1 != text2 {
		t.Errorf(`text1 != text2`)
	}
}

 //Tests to see if variable contains a value 
 func TestNotEmpty(t *testing.T){
	text4 := errors.New("Something")
   if text4 == nil{
	t.Errorf(`"text4 is nil "`)
}
 }

 //Tests to see if error message will work 
 func TestErrorMessage(t *testing.T) {
    
	message := errors.New("x and y are different values")

	x := 2
	y := 5
	if x != y {
		fmt.Print(message)
	}

	
	}










