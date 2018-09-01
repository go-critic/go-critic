package checker_test

// https://youtu.be/5DVV36uqQ4E?t=1334

type foo struct {
	k string
}

func (f foo) bar(i int) {
	println(i)
}

func main() {
	f := foo{}

	f.bar(10)

	/// consider to change to `f.bar(20)`
	foo.bar(f, 20)
}
