package checker_test

var (
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
