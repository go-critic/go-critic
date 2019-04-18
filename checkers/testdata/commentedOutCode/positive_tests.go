package checker_test

func singleLineCode() {
	/*! may want to remove commented-out code */
	//fmt.Printf("Operand: %v\n", p.input)

	/*! may want to remove commented-out code */
	// warnings := make(map[int][]*warning)

	/*! may want to remove commented-out code */
	// return &goldenFile{warnings: warnings}
}

func multiLineCode() {
	/*! may want to remove commented-out code */
	// if !strings.HasPrefix(l, "// ") {
	// }

	/*! may want to remove commented-out code */
	/*
		if false {
			log.Printf("this is forgotten debug code")
		}
	*/

	/*! may want to remove commented-out code */
	/*
		resp, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			e := err.Error()
			return c.JSON(http.StatusNotFound, structs.Response{Ok: false, Reason: &e})
		}
		fmt.Println(string(resp))
		return c.JSON(http.StatusNotFound, structs.Response{Ok: false})
	*/

	/*! may want to remove commented-out code */
	//rulebases.POST("/", postRsHandler)                                //create a rulebase
	//rulebases.DELETE("/:id", deleteRsHandler)                         //delete a rulebase
	//rulebases.PUT("/:setid", putRsHandler)                            //update a rulebase
}
