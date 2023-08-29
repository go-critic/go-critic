package checker_test

/*! put a space between `//` and comment text */
//this is a comment without leading space

/*! put a space between `//` and comment text */
//nolinto

/*! put a space between `//` and comment text */
//no inspection

/*! put a space between `//` and comment text */
//no inspection something

func f1() {
	/*! put a space between `//` and comment text */
	//block with
	//several lines
	//without leading space
}

var (
	/*! put a space between `//` and comment text */
	//only several
	//lines don't follow
	// the convention
	x = 10
)
