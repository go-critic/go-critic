package checker_test

func explicitNil() {
	_ = func(err error) error {
		if err == nil {
			return nil
		}
		return nil
	}

	_ = func(o *object) *object {
		if o == nil {
			return nil
		}
		return &object{}
	}

	_ = func(o *object) *byte {
		if o.data == nil {
			return nil
		}
		return nil
	}

	_ = func(pointers [][][]map[string]*int) *int {
		if pointers[0][1][2]["ptr"] == nil {
			return nil
		}
		if ptr := pointers[0][1][2]["ptr"]; ptr == nil {
			return nil
		}
		return nil
	}
}

func explicitNotEqual() {
	_ = func(err error) error {
		if err != nil {
			return err
		}
		return nil
	}

	_ = func(o *object) *object {
		if o != nil {
			return o
		}
		return &object{}
	}

	_ = func(o *object) *byte {
		if o.data != nil {
			return o.data
		}
		return nil
	}

	_ = func(pointers [][][]map[string]*int) *int {
		if pointers[0][1][2]["ptr"] != nil {
			return pointers[0][1][2]["ptr"]
		}
		if ptr := pointers[0][1][2]["ptr"]; ptr != nil {
			return ptr
		}
		return nil
	}
}
