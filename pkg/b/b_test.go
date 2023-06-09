// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package b_test

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
	"testing"
	
	"github.com/petermetz/go-monorepo-test/pkg/b/v2"
)

type SomethingB struct {
	
}

// Main function
func TestB(t *testing.T) {

	// Sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// Finding the index
	fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())
	doSomethingBWouldDo()
}

func doSomethingBWouldDo() {
	fmt.Println("This is something only B would do!")
	b.CallB("B")
}
