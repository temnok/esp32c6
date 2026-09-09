
#include "../lib/csr/csr.h"
#include "../lib/fmt/fmt.h"
#include "../lib/sys/sys.h"

__attribute__((section(".text._vectors"), naked, used, retain))
void _vectors() {
    __asm__ volatile (
        ".rept 32\n"
        "   j trap_handler\n"
        ".endr\n"
    );
}

const unsigned GPIO_MATRIX_BASE_ADDR = 0x60091000;

unsigned
    * const GPIO_OUT_W1TS_REG = (unsigned*)(GPIO_MATRIX_BASE_ADDR + 0x0008),
    * const GPIO_ENABLE_W1TS_REG = (unsigned*)(GPIO_MATRIX_BASE_ADDR + 0x0024),
    * const GPIO_FUNC_OUT_SEL_CFG_REG = (unsigned*)(GPIO_MATRIX_BASE_ADDR + 0x0554);

__attribute__((section(".text._start")))
void _start() {
    extern int _bss_start, _bss_end;
    for (int *p = &_bss_start; p < &_bss_end; p++) *p = 0;

//    fmt_str(sys_print, "Reg value: ");
//    fmt_unsigned_hex(sys_print, GPIO_FUNC_OUT_SEL_CFG_REG[15]);
//    fmt_str(sys_print, "\n");
    *GPIO_ENABLE_W1TS_REG = 1<<15;
    *GPIO_OUT_W1TS_REG = 1<<15;

    sys_exit();
}

__attribute__ ((interrupt, used, retain))
void trap_handler() {
    sys_exit();
}
