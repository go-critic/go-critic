package checker_test

import (
	"log"
)

func noWarnings(x bool) {
	if x {
		log.Print()
		log.Fatal()
	}
	log.Print()
	log.Fatal()
}

func noWarningsWhenReturnInDifferentBlock(x bool) {
	if x {
		log.Print()
		log.Fatal()
	}
	log.Print()
	return
}
