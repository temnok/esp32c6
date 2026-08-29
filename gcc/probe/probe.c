
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
    fmt_str(sys_print, "Hello, world!\n");

    sys_exit();
}

__attribute__ ((interrupt))
void trap_handler() {
    fmt_str(sys_print, "\nTrap: mcause=0x");
    fmt_unsigned_hex(sys_print, csr_read(csr_mcause));
    fmt_str(sys_print, ", mepc=0x");
    fmt_unsigned_hex(sys_print, csr_read(csr_mepc));
    fmt_str(sys_print, "\n");

    sys_exit();
}
