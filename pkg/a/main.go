// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package a

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type SomethingA struct {
	
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

}

func doSomethingAWouldDo() {
	fmt.Println("This is something only A would do!")
}
