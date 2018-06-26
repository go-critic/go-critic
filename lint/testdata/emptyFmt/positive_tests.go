package checker_test

import (
	"fmt"
	"log"
)

func f() {
	/// consider to change function to fmt.Sprint
	_ = fmt.Sprintf("whatever")

	/// consider to change function to fmt.Fprint
	fmt.Fprintf(nil, "whatever")

	/// consider to change function to errors.New
	_ = fmt.Errorf("whereever")

	/// consider to change function to log.Print
	log.Printf("whenever")
}
