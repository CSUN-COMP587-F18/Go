package main

//Purpose: to show that Golang will not execute invalid code 
//and that the typechecker will reject ill-typed programs


//import ("testing")

//Test case works without any problems since syntax is correct

func TestMax(t *testing.T) {
	if max(1,2) != 2 {
		t.Error("Not 2")
	}
}


//This test case won't work since the syntax for max 1,2 is not in the form max(1,2)

func TestMaxx(t *testing.T) {
	if max 1,3 != 3 {
		t.Error("Not 3")
	}
}

/*
When gofmt -e syntax_test.go is ran the following occurs:
______________________________________________________

gofmt -e syntaxcorrect_test.go
syntaxcorrect_test.go:20:9: expected ';', found 1
syntaxcorrect_test.go:20:10: expected operand, found ','
syntaxcorrect_test.go:28:21: expected ';', found 'EOF'
syntaxcorrect_test.go:28:21: expected ';', found 'EOF'
syntaxcorrect_test.go:28:21: expected '{', found 'EOF'
syntaxcorrect_test.go:28:21: expected '}', found 'EOF'
syntaxcorrect_test.go:28:21: expected '}', found 'EOF'
________________________________________________________

If there are no syntax errrors the whole file will be displayed


*/