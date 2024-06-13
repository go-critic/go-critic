package checker_test

import (
	"strconv"
	"sync"
)

type T2 struct {
	val int
}

type T1 struct {
	fun   func()
	val   int
	child T2
	items []string
	any   interface{}
}

func recv4() (int, int, int, error) {
	return 0, 0, 0, nil
}

func recv1() (int, error) {
	return 0, nil
}

func recvT1() (T1, error) {
	return T1{}, nil
}

func recvErr() error {
	return nil
}

func recvFunc() (func(), error) {
	return nil, nil
}

func send(args ...interface{}) interface{} {
	return nil
}

type custErr struct{}

func (custErr) Error() string {
	return `error`
}

func custErrFunc() error {
	return custErr{}
}

func testFunc() {

	var err error
	_ = err

	var xerr error
	_ = xerr

	/*! no error access */
	f, xerr := recvFunc()
	/*! expr, missing error check accessing [f] */
	f()

	f, xerr = recvFunc()
	if xerr != nil {

	}
	f()

	a, b, c, _ := recv4()
	send(a, b, c)

	/*! no error access */
	xerr = custErrFunc()

	/*! no error access */
	a, b, c, err = recv4()
	/*! expr, missing error check accessing [a b c] */
	send(a, b, c)

	/*! no error access */
	f, err = recvFunc()
	/*! expr, missing error check accessing [f] */
	f()

	var t1 T1
	t1.fun, err = recvFunc()
	_ = t1.val

	/*! assign, missing error check accessing [t1.fun] */
	_ = t1.fun

	/*! no error access */
	t1.val, err = recv1()
	/*! expr, missing error check accessing [t1.val] */
	send(t1.val)

	/*! no error access */
	t1.val, err = recv1()
	/*! expr, missing error check accessing [t1.val] */
	send(send(t1.val))

	/*! no error access */
	t1.child.val, err = recv1()
	/*! expr, missing error check accessing [t1.child.val] */
	send(t1.child.val)

	a, err = recv1()
	/*! if(1), missing error check accessing [a] */
	if a != 0 {
		send(a)
	}

	a, err = recv1()
	if a != 0 && err != nil {
		send(a)
	}

	a, err = recv1()
	if b == 0 {
		send(b)
	}
	send(a)

	/*! no error access */
	a, err = recv1()
	/*! expr, missing error check accessing [a] */
	send(a)

	/*! no error access */
	a, err = recv1()
	/*! no error access */
	b, err = recv1()

	t1, err = recvT1()
	/*! range, missing error check accessing [t1] */
	for _, item := range t1.items {
		send(item)
	}

	a, err = recv1()
	/*! assign, missing error check accessing [a] */
	_ = send(a)

	a, err = recv1()
	/*! if, missing error check accessing [a] */
	if b := send(a); b != nil {
		send(b)
	}

	/*! no error access */
	a, err = recv1()
	/*! go, missing error check accessing [a] */
	go send(a)

	/*! no error access */
	a, err = recv1()
	go send(err)
	/*! expr, missing error check accessing [a] */
	send(a)

	/*! no error access */
	a, err = recv1()
	go send(a, err)
	/*! expr, missing error check accessing [a] */
	send(a)

	a, err = recv1()
	if err != nil {
		send(err)
	}
	send(a)

	/*! no error access */
	a, err = recv1()
	/*! inc/dec, missing error check accessing [a] */
	a++

	ch := make(chan interface{}, 1)

	/*! no error access */
	a, err = recv1()
	/*! send, missing error check accessing [a] */
	ch <- a

	/*! no error access */
	a, err = recv1()
	/*! send, missing error check accessing [a] */
	ch <- send(a)

	a, err = recv1()
	/*! switch, missing error check accessing [a] */
	switch a {
	}

	t1, err = recvT1()
	/*! switch, missing error check accessing [t1] */
	switch t1.any.(type) {
	}
}

func testFunc2() {

	var err error
	_ = err

	if 1 == 1 {

		var a int
		/*! no error access */
		a, err = recv1()
		/*! expr, missing error check accessing [a] */
		send(a)

		a, err = recv1()
		/*! switch, missing error check accessing [a] */
		switch a {
		case 1:
			a, err = recv1()
			/*! expr, missing error check accessing [a] */
			send(a)
		}

		if a == 1 {
		} else if a == 2 {

			a, err = recv1()
			/*! expr, missing error check accessing [a] */
			send(a)
		}

		if a == 1 {
		} else {

			a, err = recv1()
			/*! switch, missing error check accessing [a] */
			switch a {
			case 1:
				a, err = recv1()
				/*! expr, missing error check accessing [a] */
				send(a)
			}
		}

		ch := make(chan interface{}, 1)
		ch <- 1

		select {
		case b := <-ch:
			_ = b
			a, err = recv1()
			/*! expr, missing error check accessing [a] */
			send(a)
		}

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			a, err = recv1()
			/*! expr, missing error check accessing [a] */
			send(a)
			wg.Done()
		}()
		wg.Wait()
	}
}

func testFunc3() {

	var err error
	_ = err

	/*! no error access */
	err = recvErr()

	/*! no error access */
	_, err = recv1()

	f, err := recvFunc()
	/*! defer, missing error check accessing [f] */
	defer f()
	if err != nil {

	}

	a, err := recv1()
	msg := err.Error() + strconv.Itoa(a)
	_ = msg
}
