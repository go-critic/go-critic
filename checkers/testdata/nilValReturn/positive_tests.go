package checker_test

type object struct {
	data *byte
}

func suspiciousReturns() {
	_ = func(err error) error {
		if err == nil {
			/*! returned expr is always nil; replace err with nil */
			return err
		}
		return nil
	}

	_ = func(o *object) *object {
		if o == nil {
			/*! returned expr is always nil; replace o with nil */
			return o
		}
		return &object{}
	}

	_ = func(o *object) *byte {
		if o.data == nil {
			/*! returned expr is always nil; replace o.data with nil */
			return o.data
		}
		return nil
	}

	_ = func(o *object, err error) (*object, error) {
		if err == nil {
			/*! returned expr is always nil; replace err with nil */
			return nil, err
		}
		return o, nil
	}

	_ = func(o *object, a *object) *object {
		if o == nil {
			/*! returned expr is always nil; replace o with nil */
			return o
		}
		return a
	}

	_ = func(pointers [][][]map[string]*int) *int {
		if pointers[0][1][2]["ptr"] == nil {
			/*! returned expr is always nil; replace pointers[0][1][2]["ptr"] with nil */
			return pointers[0][1][2]["ptr"]
		}
		if ptr := pointers[0][1][2]["ptr"]; ptr == nil {
			/*! returned expr is always nil; replace ptr with nil */
			return ptr
		}
		return nil
	}
}
