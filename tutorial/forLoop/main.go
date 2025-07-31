package main

import "fmt"

func main() {

	// normal for
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// for used as while
	sum := 1
	for sum <= 100 {
		sum += sum
	}
	fmt.Println(sum)

	// runs forever
	//for {
	//
	//}

	//

}
