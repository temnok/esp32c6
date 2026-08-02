package check

func E(err error) {
	if err != nil {
		panic(err)
	}
}

func E1[A any](a A, err error) A {
	E(err)

	return a
}

func Call(f func() error) {
	E(f())
}
