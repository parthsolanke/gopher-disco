package main

import "fmt"

// eg: closuer which captures and preserves teh state of i
// local variables which are short lived are stored in the stack
// but in case of a clusere teh i is stored on teh heap adn refrenced by the clouser
// hence the state is preserved as i will be long lived
func intSeries() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nxtInt := intSeries()

	fmt.Println(nxtInt())
	fmt.Println(nxtInt())

	// state is unique to that particular function
	nxtInt2 := intSeries()
	fmt.Println(nxtInt2())

	fmt.Println(nxtInt())
	fmt.Println(nxtInt())
	fmt.Println(nxtInt())
}
