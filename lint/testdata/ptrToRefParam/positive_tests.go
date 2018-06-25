package checker_test

/// consider `m' to be of non-pointer type
/// consider `k' to be of non-pointer type
func f1(m *map[int]string) (k *chan float64) {
	return nil
}

/// consider `ch' to be of non-pointer type
func f2(ch *chan string) {}

/// consider `m' to be of non-pointer type
func f3(a int, m *map[int]string, s string) {}

/// consider `slice' to be of non-pointer type
/// consider `ch' to be of non-pointer type
func f4(slice *[]string) (ch *chan *int) {
	return nil
}

/// consider `a' to be of non-pointer type
/// consider `b' to be of non-pointer type
/// consider to make non-pointer type for `*chan *int`
func f5(a, b *chan string) *chan *int {
	return nil
}

/// consider `a' to be of non-pointer type
/// consider `b' to be of non-pointer type
func f6(c int, a, b *chan string) {}

/// consider `a' to be of non-pointer type
/// consider `b' to be of non-pointer type
func f7() (a, b *chan string) {
	return nil, nil
}
