package checker_test

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func g1() {
	_ = fmt.Sprintf("%s", "whatever")

	fmt.Fprintf(nil, "%s", "whatever")

	_ = fmt.Errorf("%s", "whereever")

	log.Printf("%s", "whenever")

	errors.Wrapf(nil, "%s", "however")
}
