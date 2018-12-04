package checker_test

import (
	"fmt"
	"reflect"
)

func warningsEmptyFallthrough(i int) bool {
	switch i {
	case 0:
		/*! replace empty case containing only fallthrough with expression list */
		fallthrough
	case 1:
		/*! replace empty case containing only fallthrough with expression list */
		fallthrough
	case 2:
		return true
	default:
		return false
	}
}

func warningsEmptyFallthrough2(kind reflect.Kind) reflect.Kind {
	switch kind {
	case reflect.Int:
		/*! replace empty case containing only fallthrough with expression list */
		fallthrough
	case reflect.Int32:
		return reflect.Int
	}
	return reflect.Invalid
}

func warningsEmptyFallthroughToDefault(i int) bool {
	switch i {
	case 0:
		return true
	case 1:
		/*! remove empty case containing only fallthrough to default case */
		fallthrough
	case 2:
		/*! remove empty case containing only fallthrough to default case */
		fallthrough
	default:
		return false
	}
}

func warningsEmptyFallthroughToNonLastDefault(i int) bool {
	switch i {
	case 0:
		return true
	case 1:
		/*! remove empty case containing only fallthrough to default case */
		fallthrough
	case 2:
		/*! remove empty case containing only fallthrough to default case */
		fallthrough
	default:
		return false
	case 3:
		return true
	}
}

func warningsNestedSwitchMixedFallthroughs(i, j int) bool {
	switch i {
	case 0:
		/*! replace empty case containing only fallthrough with expression list */
		fallthrough
	case 1:
		switch j {
		case 0:
			/*! replace empty case containing only fallthrough with expression list */
			fallthrough
		case 1:
			fmt.Println("")
			fallthrough
		case 2:
			return true
		case 3:
			/*! remove empty case containing only fallthrough to default case */
			fallthrough
		default:
			return false
		}
	case 2:
		return true
	case 3:
		/*! remove empty case containing only fallthrough to default case */
		fallthrough
	default:
		return false
	}
}
