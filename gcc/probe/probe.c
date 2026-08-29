
#include "../lib/csr/csr.h"
#include "../lib/fmt/fmt.h"
#include "../lib/sys/sys.h"

__attribute__((naked, section(".text._vectors")))
void _vectors() {
    __asm__ volatile (
        "j trap_handler\n" //  0
        "j trap_handler\n" //  1
        "j trap_handler\n" //  2
        "j trap_handler\n" //  3
        "j trap_handler\n" //  4
        "j trap_handler\n" //  5
        "j trap_handler\n" //  6
        "j trap_handler\n" //  7
        "j trap_handler\n" //  8
        "j trap_handler\n" //  9
        "j trap_handler\n" // 10
        "j trap_handler\n" // 11
        "j trap_handler\n" // 12
        "j trap_handler\n" // 13
        "j trap_handler\n" // 14
        "j trap_handler\n" // 15
        "j trap_handler\n" // 16
        "j trap_handler\n" // 17
        "j trap_handler\n" // 18
        "j trap_handler\n" // 19
        "j trap_handler\n" // 20
        "j trap_handler\n" // 21
        "j trap_handler\n" // 22
        "j trap_handler\n" // 23
        "j trap_handler\n" // 24
        "j trap_handler\n" // 25
        "j trap_handler\n" // 26
        "j trap_handler\n" // 27
        "j trap_handler\n" // 28
        "j trap_handler\n" // 29
        "j trap_handler\n" // 30
        "j trap_handler\n" // 31
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
