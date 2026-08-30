
#include "../lib/csr/csr.h"
#include "../lib/fmt/fmt.h"
#include "../lib/sys/sys.h"

__attribute__((naked, section(".text._vectors")))
void _vectors() {
    __asm__ volatile (
        ".rept 32\n"
        "   j trap_handler\n"
        ".endr\n"
    );
}

__attribute__((section(".text._start")))
void _start() {
    sys_exit();
}

__attribute__ ((interrupt))
void trap_handler() {
    sys_exit();
}
