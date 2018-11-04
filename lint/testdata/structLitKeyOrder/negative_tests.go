package checker_test

var (
	_ = foo{pt: point{}, b: 1}
	_ = foo{bar: bar{}, c: 1}
	_ = foo{a: 1, pt: point{x: 2, y: 1}}

	_ = point{}

	_ = point{x: 0}
	_ = point{y: 1}
	_ = point{z: 2}

	_ = point{x: 0, y: 1}
	_ = point{x: 0, z: 2}
	_ = point{y: 0, z: 1}

	_ = point{x: 0, y: 1, z: 2}
)

func consistentKeysOrder() {
	_ = point{}

	_ = point{x: 0}
	_ = point{y: 1}
	_ = point{z: 2}

	_ = point{x: 0, y: 1}
	_ = point{x: 0, z: 2}
	_ = point{y: 0, z: 1}

	_ = point{x: 0, y: 1, z: 2}
}
