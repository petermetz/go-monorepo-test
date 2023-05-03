package a

// Importing packages
import (
	"fmt"
)

func CallA(caller string) {
	fmt.Printf("A: Called A from %s\n", caller)
	fmt.Println("A: This is something only A would do!")
	fmt.Println("A: End")
}