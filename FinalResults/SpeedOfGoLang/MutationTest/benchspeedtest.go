package main

//import ("testing")

//sorts array of 500 random values
func BenchmarkTest500(/*b *testing.B*/) {
    performSort(500) 
}

//sorts array of 1000 random values
func BenchmarkTest1000(/*b *testing.B*/) {
    performSort(1000)
}

//sorts array of 500000 random values
func BenchmarkTest500000(/*b *testing.B*/) {
    performSort(500000)
}

//sorts array of 1000000 random values
func BenchmarkTest1Mill(/*b *testing.B*/) {
    performSort(1000000)
}

//sorts array of 500000000 random values
func BenchmarkTestFiveHundredMill(/*b *testing.B*/) {
    performSort(500000000)
}

/*
$ go test -bench=.
goos: windows
goarch: amd64
BenchmarkTest500-4                      2000000000               0.00 ns/op
BenchmarkTest1000-4                     2000000000               0.00 ns/op
BenchmarkTest500000-4                   1000000000               0.14 ns/op
BenchmarkTest1Mill-4                    1000000000               0.24 ns/op
BenchmarkTestFiveHundredMill-4                 1        157220687700 ns/op
PASS
ok    163.190s

*/