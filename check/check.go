package check

func A(err error) {
	if err != nil {
		panic(err)
	}
}

func B[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}

	return a
}

func Defer(f func() error) {
	A(f())
}
