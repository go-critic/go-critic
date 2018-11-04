package checker_test

import "time"

func foo() {}

func bar() {
	time.Sleep(time.Hour)
}

var ch = make(chan int)

func g1() {
	for i := 0; i < 10; i++ {
		select {
		case <-time.Tick(1 * time.Second):
		default:
		}
	}
}

func g2() {
	for range ch {
		select {
		case <-time.Tick(1 * time.Second):
		default:
			foo()
		}
	}
}

func g3() {
	for {
		select {
		case <-ch:
		default:
			ch <- 10
		}
	}
}

func g4() {
	for {
		select {
		default:
			time.Sleep(10 * time.Second)
		}
	}
}

func g5() {
	for {
		select {
		case <-ch:
		default:
			<-time.After(10 * time.Hour)
		}
	}
}

func g6() {
	for {
		select {
		case <-ch:
		default:
			<-time.Tick(10 * time.Hour)
		}
	}
}

func g7() {
	done := make(chan int)
	msgs := make(chan int)
	m := 10
	for {
		msgs <- m
		select {
		case <-done:
			return
		default:
		}
	}
}
