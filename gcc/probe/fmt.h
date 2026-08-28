#pragma once

void fmt_unsigned(void print(char), unsigned val) {
    static unsigned pow[] = {1000000000, 100000000, 10000000, 1000000, 100000, 10000, 1000, 100, 10, 1};

    int i = 0, n = 10;
    while (pow[i] > val && i+1 < n) {
        i++;
    }

    for (; i < n; i++) {
        char d = '0';

        unsigned p = pow[i], p2 = p+p, p4 = p2+p2;
        if (val >= p4) { d += 4; val -= p4; }
        if (val >= p2) { d += 2; val -= p2; }
        if (val >= p2) { d += 2; val -= p2; }
        if (val >= p) { d++; val -= p; }

        print(d);
    }
}

void fmt_int(void print(char), int val) {
    if (val < 0) {
        print('-');
        val = -val;
    }

    fmt_unsigned(print, (unsigned)val);
}
