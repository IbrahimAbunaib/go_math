package calculate

import "fmt"

func Add(a, b int) int {
	return a + b // returns the sum
}

func main() {
	a, b := 4, 9
	fmt.Println("the sum of both parameters are", Add(a, b))
}
