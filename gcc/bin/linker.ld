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

        ASSERT(. == 0x80, "error: vector table size must be 0x80");
        KEEP(*(.text._start))
        *(.text .text.*)
    } > RAM

    .rodata :
    {
        . = ALIGN(4);
        *(.srodata .srodata.*)

        . = ALIGN(4);
        *(.rodata .rodata.*)
    } > RAM

    .data :
    {
        . = ALIGN(4);
        *(.sdata .sdata.*)

        . = ALIGN(4);
        *(.data .data.*)
    } > RAM

    .bss :
    {
        . = ALIGN(4);
        PROVIDE(_bss_start = .);
        *(.sbss .sbss.*)

        . = ALIGN(4);
        *(.bss .bss.*)

        . = ALIGN(4);
        PROVIDE(_bss_end = .);
    } > RAM
}
