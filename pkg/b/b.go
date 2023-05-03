package b

// Importing packages
import (
	"fmt"
	
	"github.com/petermetz/go-monorepo-test/pkg/a"
)

func CallB(caller string) {
	fmt.Printf("B: Called B from %s\n", caller)
	fmt.Println("B: This is something only B would do!")
	fmt.Println("B: Calling A")
	a.CallA("B")
	fmt.Println("B: End")
}