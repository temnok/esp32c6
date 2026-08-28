
#include "fmt.h"

__attribute__((section(".text._start")))
void _start() {
    char str[10], *pos = str, *limit = str + sizeof(str);

    void append(char c) {
        if (pos < limit) {
            *pos++ = c;
        }
    }

    fmt_int(append, -123456789);

    append('\0');

    __builtin_trap();
}
