package asm

func (asm *asm) DIV(rd, rs1, rs2 int)    { asm.R(0x02004033, rd, rs1, rs2) }
func (asm *asm) DIVU(rd, rs1, rs2 int)   { asm.R(0x02005033, rd, rs1, rs2) }
func (asm *asm) MUL(rd, rs1, rs2 int)    { asm.R(0x02000033, rd, rs1, rs2) }
func (asm *asm) MULH(rd, rs1, rs2 int)   { asm.R(0x02001033, rd, rs1, rs2) }
func (asm *asm) MULHSU(rd, rs1, rs2 int) { asm.R(0x02002033, rd, rs1, rs2) }
func (asm *asm) MULHU(rd, rs1, rs2 int)  { asm.R(0x02003033, rd, rs1, rs2) }
func (asm *asm) REM(rd, rs1, rs2 int)    { asm.R(0x02006033, rd, rs1, rs2) }
func (asm *asm) REMU(rd, rs1, rs2 int)   { asm.R(0x02007033, rd, rs1, rs2) }
