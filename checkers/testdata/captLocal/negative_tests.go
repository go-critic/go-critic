package checker_test

func goodFunc1(in int) (out int) {
	return 0
}

func goodFunc2(in, x int) (y, z int) {
	return 0, 0
}

type emptyStruct struct{}

func (in emptyStruct) method1(out *int) {}

type ExportedType int

func (x ExportedType) capitalizedType(y ExportedType) ExportedType {
	return 0
}

func noWarnings() {
	type LocalCapitalizedType int
	type (
		T1 float32
		T2 float64
	)

	v1, v2 := 1, 2
	var x, y = v1, v2

	{
		v3, v4 := x, y
		v5, v3 := v4, v3
		_, _ = v3, v5

		const (
			c1         = 1
			c2         = 2
			c3, c4, c5 = 3, 4, 5
		)
	}
}
