package check

import "runtime/debug"

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

func RecoverAndPrintStack(print func(...any)) {
	if err := recover(); err != nil {
		print(err, "\n", string(debug.Stack()), "\n")
	}
}
