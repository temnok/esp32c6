package isa

type RV32IMAC interface { // 86 =
	RV32I //                 40 +
	RV32M //                  8 +
	RV32A //                 11 +
	RV32C //                 27
}

type RV32IMACpcf interface { // 98 =
	RV32IMAC   //               86 +
	Privileged //                5 +
	Zicsr      //                6 +
	Zifencei   //                1
}
