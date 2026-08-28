#pragma once

__attribute__((always_inline))
static inline int csr_read(int csr) {
    int val;
    __asm__ __volatile__ ("csrr %0, %1" : "=r"(val) : "n"(csr));
    return val;
}

__attribute__((always_inline))
static inline void csr_write(int csr, int val) {
    __asm__ __volatile__ ("csrw %0, %1" : : "n"(csr), "r"(val));
}

__attribute__((always_inline))
static inline void csr_write_imm(int csr, int val) {
    __asm__ __volatile__ ("csrwi %0, %1" : : "n"(csr), "n"(val));
}

__attribute__((always_inline))
static inline void csr_clear(int csr, int val) {
    __asm__ __volatile__ ("csrc %0, %1" : : "n"(csr), "r"(val));
}

__attribute__((always_inline))
static inline void csr_clear_imm(int csr, int val) {
    __asm__ __volatile__ ("csrci %0, %1" : : "n"(csr), "n"(val));
}

__attribute__((always_inline))
static inline void csr_set(int csr, int val) {
    __asm__ __volatile__ ("csrs %0, %1" : : "n"(csr), "r"(val));
}

__attribute__((always_inline))
static inline void csr_set_imm(int csr, int val) {
    __asm__ __volatile__ ("csrsi %0, %1" : : "n"(csr), "n"(val));
}
