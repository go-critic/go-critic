package checker_test

import (
	"time"
)

func f1() {
	ch := make(chan int)
	for {
		select {
		/// default case without a blocking operation or sleep might waste a CPU time
		default:
		case <-ch:
		}
	}
}

func f2() {
	for {
		select {
		case <-time.Tick(1 * time.Second):
		/// default case without a blocking operation or sleep might waste a CPU time
		default:
			foo()
		}
	}
}

func f4() {
	for i := 0; ; i++ {
		select {
		case <-time.Tick(1 * time.Second):
			/// default case without a blocking operation or sleep might waste a CPU time
		default:
		}
	}
}

func f5() {
	for {
		select {
		case <-ch:
			/// default case without a blocking operation or sleep might waste a CPU time
		default:
			bar()
		}
	}
}
