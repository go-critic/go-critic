package checker_test

func methodCalls() {
	f := foo{}
	f.bar(20)
	f.bar2(20, "str")
	f.ptrRecv()
}

func nilCall() {
	iface.ptrRecv(nil)
	(*foo).ptrRecv(nil)
}
