#pragma once

void sys_exit() {
    __builtin_trap();
}
