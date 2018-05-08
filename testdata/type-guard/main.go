package linter_test

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
	switch v.(type) {
	case int:
	case float32:
	}
	switch v.(type) {
	}
}

func typeGuard4()

func typeGuard5(v interface{}) int {
	switch v.(type) {
	case int:
		return v.(int)
	case float32, float64:
		switch v.(type) {
		case float32:
			return int(v.(float32))
		case float64:
			return int(v.(float64))
		}
	default:
		switch v.(type) {
		case int32:
			return int(v.(int32))
		}
	}
	return 0
}
