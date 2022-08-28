package checker_test

const length = 1

var opindex [(length + 1) & 3]*int

var _ [(length + 100 - 20*5)]*int

var _ func(int, string)

type myString string

func convertPtr(x string) *myString {
	return (*myString)(&x)
}

func multipleReturn() (int, bool) {
	return 1, true
}

type goodMap1 map[string]string

type goodMap2 map[[5][5]string]map[string]string

var _ = [4]*int{}

var _ = func() []func() { return nil }

var f52want float64 = 1.0 / (1 << 52)

const (
	c1 = (1 + (2 + 3))
	c2 = (1 << 2 << 3)
)

func f() {
	const (
		localC1 = (1 + (2 + 3))
		localC2 = (1 << 2 << 3)
	)

}

type ifaceToEmbed interface{}

type ifaceWithEmbedding interface {
	ifaceToEmbed
	Foo()
}

type structToEmbed struct{}

type structWithEmbedding struct {
	structToEmbed
	Field int
}

var _ interface{} = (*int)(nil)

type myWriter interface {
	myWrite()
}

type noopWriter struct{}

func (*noopWriter) myWrite() {}

var _ myWriter = (*noopWriter)(nil)

func ptrCast() {
	_ = (*int)(nil)

	_ = (*int)((*int)(nil))

	_ = (***int)(nil)
}

func goodMethodExpr() {
	_ = (*noopWriter).myWrite
}

func channelIssue1035() {
	type WebsocketMsg struct{}

	_ = make(chan (<-chan *WebsocketMsg))
	_ = make(<-chan int)
	_ = make(chan<- int)
	_ = make(chan int)
	_ = make(chan int)

	var _ chan (<-chan *WebsocketMsg)
	var _ chan (<-chan int)
	var _ chan (chan<- int)
	var _ chan<- (chan<- int)
	var _ <-chan (chan<- int)
	var _ chan (<-chan (chan int))
	var _ chan (<-chan (<-chan int))
}

func funcIssue1245() {
	_ = (func())(nil)
}
