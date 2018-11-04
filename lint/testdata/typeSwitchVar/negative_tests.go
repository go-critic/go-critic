package checker_test

type negativeTypeSwitchVarPoint struct {
	x int
	y int
}

func negativeTypeSwitchVarF1() int {
	var v interface{} = negativeTypeSwitchVarPoint{1, 2}

	switch v := v.(type) {
	case int:
		return v
	case negativeTypeSwitchVarPoint:
		return v.x + v.y
	default:
		return 0
	}
}

func negativeTypeSwitchVarF2() int {
	xs := [][]interface{}{
		{1, 2, 3},
	}

	switch xs := xs[0][0].(type) {
	default:
		return 0
	case []int:
		return xs[0]
	}
}

func negativeTypeSwitchVarF3() int {
	type nested struct {
		a struct {
			b struct {
				value interface{}
			}
		}
	}
	var v nested
	v.a.b.value = 10

	switch value := v.a.b.value.(type) {
	case int8, int16:
		return 16
	case int32:
		return 32
	case int:
		return value
	}
	return 0
}

func negativeTypeSwitchVarF4(x, y interface{}) int {
	switch x.(type) {
	case int:
		switch y := y.(type) {
		case int:
			// shadows outer x, so checker should no trigger.
			x := interface{}(1)
			return x.(int) + y
		}
	case float32, float64:
		switch x := x.(type) {
		case float32:
			return int(x)
		case float64:
			return int(x)
		}
	default:
		switch x := x.(type) {
		case int32:
			return int(x)
		}
	}
	return 0
}

func negativeTypeSwitchVarF5(x, y, z interface{}) int {
	switch x := x.(type) {
	case int:
		switch y := y.(type) {
		case int:
			switch z := z.(type) {
			case int:
				return x + y + z
			}
		}
	}
	return 0
}
