
#include "../lib/csr/csr.h"
#include "../lib/fmt/fmt.h"
#include "../lib/sys/sys.h"
#include "../lib/fixed/fixed.h"

__attribute__((naked, section(".text._vectors")))
void _vectors() {
    __asm__ volatile (
        ".rept 32\n"
        "   j trap_handler\n"
        ".endr\n"
    );
}

static int tmp;

__attribute__((section(".text._start")))
void _start() {
    *(int*)0x40870000 = fixed_mul(*(int*)0x40870004, *(int*)0x40870008);

    sys_exit();
}

__attribute__ ((interrupt))
void trap_handler() {
    sys_exit();
}
