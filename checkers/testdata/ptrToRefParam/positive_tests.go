package checker_test

/*! consider `m' to be of non-pointer type */
/*! consider `k' to be of non-pointer type */
func f1(m *map[int]string) (k *chan float64) {
	return nil
}

/*! consider `ch' to be of non-pointer type */
func f2(ch *chan string) {}

/*! consider `m' to be of non-pointer type */
func f3(a int, m *map[int]string, s string) {}

/*! consider `ch' to be of non-pointer type */
func f4(slice *[]string) (ch *chan *int) {
	return nil
}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
/*! consider to make non-pointer type for `*chan *int` */
func f5(a, b *chan string) *chan *int {
	return nil
}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
func f6(c int, a, b *chan string) {}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
func f7() (a, b *chan string) {
	return nil, nil
}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
func f8(a, b *interface{}) {}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
func f9() (a, b *interface{}) {
	return nil, nil
}

type myInterface interface {
	f()
}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
func f10(a, b *myInterface) {}

/*! consider `a' to be of non-pointer type */
/*! consider `b' to be of non-pointer type */
func f11() (a, b *myInterface) {
	return nil, nil
}

type iface myInterface

/*! consider to make non-pointer type for `*iface` */
func underlyingIface(*iface) {}
