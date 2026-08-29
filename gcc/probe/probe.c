
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

void f(int *p) {
    fmt_int(sys_print, *p);
}

__attribute__((section(".text._start")))
void _start() {
    fmt_str(sys_print, "Hello, world!\n");

    sys_exit();
}

__attribute__ ((interrupt))
void trap_handler() {
    unsigned cause = csr_read(csr_mcause);

    fmt_str(sys_print, "\n\nTrap, cause 0x");
    fmt_unsigned_hex(sys_print, cause);
    fmt_str(sys_print, "\n");

    sys_exit();
}
