#pragma once

static inline int fixed_mul(int a, int b) {
    int tmp;

    __asm__ __volatile__ (
        "mul  %2, %0, %1\n"
        "mulh %0, %0, %1\n"
        "srl  %1, %2, 16\n"
        "sll  %0, %0, 16\n"
        "or  %0, %0, %1"

        : "+r" (a), "+r" (b), "=&r" (tmp)
    );

    return a;
}
