// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package c

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
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

}

func doSomethingCWouldDo() {
	fmt.Println("This is something only C would do!")
}
