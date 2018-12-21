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
