package main
import (
	"fmt"
	// "time"
)

func main() {
	if num := 10; num % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	i := 2

	fmt.Println("Write ", i, " as ")
	switch i {
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("unknown")
	}
}
