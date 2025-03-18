package checker_test

func noParams() {}

func outputParams() bigStruct {
	return bigStruct{}
}

func bigArrayPtr1(a *[200]int) {}

func bigArrayPtr2(a, b *[1024]byte) {}

func bigStructPtr1(x *bigStruct) {}

func bigStructPtr2(x, y *bigStruct) {}

func mixedBigObjectsPtr(x *bigStruct, y *[20][]int) {}

func (x *bigStruct) bigRecvPtr(y *[2]bigStruct) {}

// String must be ignored #1168.
func (x bigStruct) String() string { return "" }

func genericFunc[T comparable](x T) {}

type Maybe[V any] struct {
	value V
}

func (m Maybe[V]) Fn() {}
