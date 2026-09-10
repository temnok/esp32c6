
#include "../lib/csr/csr.h"
#include "../lib/dev/dev.h"
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

void ws2812b_write(int pin, unsigned char r, unsigned char g, unsigned char b) {
    int rgb = g<<16 | r<<8 | b;
    for (int i = 23; i >= 0; i--) {
        int bit = (rgb>>i)&1;

        *(volatile int*)GPIO_OUT_W1TS_REG = 1<<pin;
        sleep_cycles(bit? 100 : 30);

        *(volatile int*)GPIO_OUT_W1TC_REG = 1<<pin;
        sleep_cycles(bit? 100 : 170);
    }
}

__attribute__((section(".text._start")))
void _start() {
    extern int _bss_start, _bss_end;
    for (int *p = &_bss_start; p < &_bss_end; p++) *p = 0;

    ((volatile int*)GPIO_FUNC_OUT_SEL_CFG_REG)[8] = 0x80;
    *(volatile int*)GPIO_ENABLE_W1TS_REG = 1<<8;
    ws2812b_write(8, 0x00, 0x02, 0x00);

    sleep_cycles(1*160'000'000);

    sys_exit();
}

__attribute__ ((interrupt, used, retain))
void trap_handler() {
    sys_exit();
}
