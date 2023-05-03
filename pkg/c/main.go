// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package main

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
	
	"github.com/petermetz/go-monorepo-test/pkg/b"
	"github.com/petermetz/go-monorepo-test/pkg/a"
)

type SomethingC struct {
	
}

// Main function
func main() {

	// Sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// Finding the index
	fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())

	doSomethingCWouldDo()
}

func doSomethingCWouldDo() {
	fmt.Println("C: This is something only C would do!")
	fmt.Println("C: Calling A")
	a.CallA("C")
	fmt.Println("C: Calling B")
	b.CallB("C")
	fmt.Println("C: End")
}
