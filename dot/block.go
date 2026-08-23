package dot

func Block(block func(dot *Dot)) []byte {
	dot := &Dot{
		labelAddr: map[string]int{},
	}

	for {
		dot.curAddr = 0
		dot.code = dot.code[:0]
		dot.retry = false

		block(dot)

		if !dot.retry {
			break
		}
	}

	return dot.code
}
