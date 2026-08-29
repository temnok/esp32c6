#pragma once

void fmt_str(void print(char), char str[]) {
    for (int i = 0; str[i] != 0; i++) {
        print(str[i]);
    }
}

void fmt_unsigned(void print(char), unsigned val) {
    if (val < 10) {
        print('0' + val);
        return;
    }

    static unsigned pow[] = {1000000000, 100000000, 10000000, 1000000, 100000, 10000, 1000, 100, 10, 1};

    int i = 0, n = 10;
    while (i+1 < n && pow[i] > val) {
        i++;
    }

    for (; i < n; i++) {
        char d = '0';

        unsigned p1 = pow[i], p2 = p1+p1, p4 = p2+p2;
        if (val >= p4) { d += 4; val -= p4; }
        if (val >= p2) { d += 2; val -= p2; }
        if (val >= p2) { d += 2; val -= p2; }
        if (val >= p1) { d += 1; val -= p1; }

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

void fmt_unsigned_hex(void print(char), unsigned val) {
    int i = 28;
    while (i > 0 && (val>>i)&0xF == 0) {
        i -= 4;
    }

    for (; i >= 0; i -= 4) {
        char d = (val>>i)&0xF;

        if (d < 10) {
            d += '0';
        } else {
            d += 'a' - 10;
        }

        print(d);
    }
}
