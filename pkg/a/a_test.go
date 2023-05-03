// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package a_test

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
	"testing"
	
	"github.com/petermetz/go-monorepo-test/pkg/a"
)

type SomethingA struct {

}

// Main function
func TestA(t *testing.T) {

	// Sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// Finding the index
	fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())

	doSomethingAWouldDo()
}

func doSomethingAWouldDo() {
	fmt.Println("This is something only A would do!")
	a.CallA("A")
}
