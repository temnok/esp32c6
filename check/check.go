package check

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func Err1[A any](a A, err error) A {
	Err(err)

	return a
}

func Call(f func() error) {
	Err(f())
}
