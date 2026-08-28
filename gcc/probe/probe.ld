MEMORY
{
    RAM   (rwx) : ORIGIN = 0x40800000, LENGTH = 512K
}

ENTRY(_start)

SECTIONS
{
    .text :
    {
        *(.text._start)
        *(.text)
    } > RAM
}
