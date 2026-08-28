
#include "csr.h"

__attribute__((section(".text._start")))
void _start() {
    csr_write_imm(0xB00, 0x11);
    csr_write(0xB00, 0x22);
    *(int*)0x40801000 = csr_read(0xB00);

    __builtin_trap();
}
