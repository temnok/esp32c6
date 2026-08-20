package isa

type RV32IMAC interface { // 88 =
	RV32I //                 42 +
	RV32M //                  8 +
	RV32A //                 11 +
	RV32C //                 27
}

type RV32IMACNZicsrZifencei interface { // 96 =
	RV32IMAC //                            88 +
	RV32N    //                             1 +
	Zicsr    //                             6 +
	Zifencei //                             1
}
