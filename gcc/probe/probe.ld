MEMORY
{
    RAM   (rwx) : ORIGIN = 0x40840000, LENGTH = 256K
}

ENTRY(_start)

SECTIONS
{
    .text :
    {
        KEEP(*(.text._vectors))
        KEEP(*(.text._start))
        *(.text*)
    } > RAM

    .rodata :
    {
        *(.rodata*)
    } > RAM

    .data :
    {
        *(.data*)
    } > RAM
}
