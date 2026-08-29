#pragma once

register char *__sys_print_tp asm ("tp");
register char *__sys_print_sp asm ("sp");

void sys_print(char c) {
    if (__sys_print_tp < __sys_print_sp) {
        *__sys_print_tp++ = c;
    }
}
