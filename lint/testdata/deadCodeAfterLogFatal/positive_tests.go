package checker_tests

import "log"

func warningsFatal(x bool) {
	if x {
		/// remove dead code after 'log.Fatal'
		log.Fatal()
		log.Print()
	}
	/// remove dead code after 'log.Fatal'
	log.Fatal()
	return
}

func warningsFatalf(x bool) {
	/// remove dead code after 'log.Fatalf'
	log.Fatalf("")
	return
}

func warningsFatalln(x bool) {
	/// remove dead code after 'log.Fatalln'
	log.Fatalln()
	return
}

func warningsPanic(x bool) {
	/// remove dead code after 'log.Panic'
	log.Panic()
	return
}

func warningsPanicf(x bool) {
	/// remove dead code after 'log.Panicf'
	log.Panicf("")
	return
}

func warningsPanicln(x bool) {
	/// remove dead code after 'log.Panicln'
	log.Panicln()
	return
}
