package main

import "math"
import "os"
import "strconv"
import "fmt"
func main() {
	l := os.Args[1]
	d := os.Args[2]
	s,err := strconv.ParseFloat(l, 64)
	if err != nil {
		panic(err)
	}
	d := 2 * math.Sqrt(s * math.Pi)
	l := math.Pi * d
	fmt.Println(d)
	fmt.Println(l)


}