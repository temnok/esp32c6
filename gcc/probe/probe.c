
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

        *(volatile int*)GPIO_OUT_REG = 1<<pin;
        sleep_cycles(bit? 75 : 25);

        *(volatile int*)GPIO_OUT_REG = 0<<pin;
        sleep_cycles(bit? 75 : 125);
    }
}

void rainbow(int pin, int pause) {
    for (int i = 0; i < 255; i++) {
        ws2812b_write(pin, 255-i, i, 0);
        sleep_cycles(pause);
    }

    for (int i = 0; i < 255; i++) {
        ws2812b_write(pin, 0, 255-i, i);
        sleep_cycles(pause);
    }

    for (int i = 0; i < 255; i++) {
        ws2812b_write(pin, i, 0, 255-i);
        sleep_cycles(pause);
    }
}

__attribute__((section(".text._start")))
void _start() {
    extern int _bss_start, _bss_end;
    for (int *p = &_bss_start; p < &_bss_end; p++) *p = 0;

    ((volatile int*)GPIO_FUNC_OUT_SEL_CFG_REG)[8] = 0x80;
    *(volatile int*)GPIO_ENABLE_W1TS_REG = 1<<8;

    for (int i = 0; i < 1; i++) {
        rainbow(8, 1'600'000);
    }

    sys_exit();
}

__attribute__ ((interrupt, used, retain))
void trap_handler() {
    sys_exit();
}
