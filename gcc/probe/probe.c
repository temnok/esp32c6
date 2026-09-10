
#include "../lib/c6/gpio.h"
#include "../lib/c6/rtc_wdt.h"
#include "../lib/c6/timg.h"
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

void sleep_cycles(unsigned cycles) {
    unsigned start = csr_read(CSR_MPCCR);
    while ((unsigned)csr_read(CSR_MPCCR) - start < cycles);
}

__attribute__((section(".text._start")))
void _start() {
    extern int _bss_start, _bss_end;
    for (int *p = &_bss_start; p < &_bss_end; p++) *p = 0;

    *(int*)GPIO_ENABLE_W1TS_REG = 1<<15;

    for (auto i = 0; i < 3; i++) {
        *(int*)GPIO_OUT_W1TS_REG = 1<<15;
        sleep_cycles(80'000'000);

        *(int*)GPIO_OUT_W1TC_REG = 1<<15;
        sleep_cycles(80'000'000);
    }

    sys_exit();
}

__attribute__ ((interrupt, used, retain))
void trap_handler() {
    sys_exit();
}
