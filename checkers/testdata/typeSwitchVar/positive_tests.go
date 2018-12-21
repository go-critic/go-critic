package checker_test

type point struct {
	x int
	y int
}

func f1() int {
	var v interface{} = point{1, 2}

	/*! case 0 can benefit from type switch with assignment */
	/*! case 1 can benefit from type switch with assignment */
	switch v.(type) {
	case int:
		return v.(int)
	case point:
		return v.(point).x + v.(point).y
	default:
		return 0
	}
}

func f2() int {
	xs := [][]interface{}{
		{1, 2, 3},
	}

	/*! case 1 can benefit from type switch with assignment */
	switch xs[0][0].(type) {
	default:
		return 0
	case []int:
		return xs[0][0].([]int)[0]
	}
}

func f3() int {
	type nested struct {
		a struct {
			b struct {
				value interface{}
			}
		}
	}
	var v nested
	v.a.b.value = 10

	/*! case 2 can benefit from type switch with assignment */
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

func f4(x, y interface{}) int {
	switch x.(type) {
	case int:
		/*! case 0 can benefit from type switch with assignment */
		switch y.(type) {
		case int:
			// shadows outer x, so checker should not trigger.
			x := interface{}(1)
			return x.(int) + y.(int)
		}
	case float32, float64:
		/*! case 0 can benefit from type switch with assignment */
		/*! case 1 can benefit from type switch with assignment */
		switch x.(type) {
		case float32:
			return int(x.(float32))
		case float64:
			return int(x.(float64))
		}
	default:
		/*! case 0 can benefit from type switch with assignment */
		switch x.(type) {
		case int32:
			return int(x.(int32))
		}
	}
	return 0
}

func f5(x, y, z interface{}) int {
	/*! case 0 can benefit from type switch with assignment */
	switch x.(type) {
	case int:
		/*! case 0 can benefit from type switch with assignment */
		switch y.(type) {
		case int:
			/*! case 0 can benefit from type switch with assignment */
			switch z.(type) {
			case int:
				return x.(int) + y.(int) + z.(int)
			}
		}
	}
	return 0
}
