
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

extern int _bss_start, _bss_end;

int static_var;

__attribute__((section(".text._start")))
void _start() {
    for (int *p = &_bss_start; p < &_bss_end; p++) {
        *p = 0;
    }

    fmt_str(sys_print, "Hello, world: ");
    fmt_int(sys_print, static_var);
    fmt_str(sys_print, "\n");

    sys_exit();
}

__attribute__ ((interrupt))
void trap_handler() {
    sys_exit();
}
