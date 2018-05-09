package checker_test

type point struct {
	x int
	y int
}

func typeGuard0() int {
	var v interface{} = point{1, 2}

	switch v.(type) {
	case int:
		return v.(int)
	case point:
		return v.(point).x + v.(point).y
	default:
		return 0
	}
}

func typeGuard1() int {
	xs := [][]interface{}{
		{1, 2, 3},
	}

	switch xs[0][0].(type) {
	default:
		return 0
	case []int:
		return xs[0][0].([]int)[0]
	}
}

func typeGuard2() int {
	type nested struct {
		a struct {
			b struct {
				value interface{}
			}
		}
	}
	var v nested
	v.a.b.value = 10

	switch v.a.b.value.(type) {
	case int8, int16:
		return 16
	case int32:
		return 32
	case int:
		return v.a.b.value.(int)
	}
	return 0
}

func typeGuard3(v interface{}) {
	// Make sure that empty switches and case clauses do not crash the checker.
	switch v.(type) {
	case int:
	case float32:
	}
	switch v.(type) {
	}
}

// Make sure that extern functions do not crash the checker.
func typeGuard4()

func typeGuard5(x, y interface{}) int {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int:
			// x shadows outer x, so checker should not trigger.
			x := interface{}(1)
			return x.(int) + y.(int)
		}
	case float32, float64:
		switch x.(type) {
		case float32:
			return int(x.(float32))
		case float64:
			return int(x.(float64))
		}
	default:
		switch x.(type) {
		case int32:
			return int(x.(int32))
		}
	}
	return 0
}

func typeGuard6(x, y, z interface{}) int {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int:
			switch z.(type) {
			case int:
				return x.(int) + y.(int) + z.(int)
			}
		}
	}
	return 0
}
