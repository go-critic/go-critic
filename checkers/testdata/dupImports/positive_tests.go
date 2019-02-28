package checker_test

/*! package is imported 3 times under different aliases on lines 4, 8 and 10 */
import printing "fmt"

import (
	/*! package is imported 3 times under different aliases on lines 4, 8 and 10 */
	"fmt"
	/*! package is imported 3 times under different aliases on lines 4, 8 and 10 */
	print "fmt"
)

func positiveHelloworld() {
	fmt.Println("Hello")
	print.Println("Shiny")
	printing.Println("World")
}
