package main

// Will need to come back and implement a multiplication, but I am going to simulate it with addition for now
import "github.com/bbetov/euler/shared"

func main() {
	i := shared.Integer{}
	i.Set("2")
	for q := 1; q < 1000; q++ {
		i.Add(&i)
	}
	println(i.SumDigits())
}
