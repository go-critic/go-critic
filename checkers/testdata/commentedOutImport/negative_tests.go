package checker_test

import (
	"errors"
)

// Check that it doesn't try to check comments after the imports spec.

var (
	//"fmt"
	_ = errors.New
)

// "fmt"
//"fmt"

/*"fmt"*/
/* "fmt" */
