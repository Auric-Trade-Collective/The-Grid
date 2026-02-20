section .text
global _start

_start:
    mov rax, 0x0a7972676e756820
    push rax
    movabs rax, 0x5d27391010101010
    add rax, 0x100010001039
    push rax

    push 0x050f
    pop rax
    mov [rel .patch + 1], ax

    push 1
    pop rax
    push rax
    pop rdi
    mov rsi, rsp
    add rsi, 2
    push 10
    pop rdx
    inc rdx
.patch:
    dw 0x0000

    xor eax, eax
    mov al, 60
    xor edi, edi
    db 0x0f, 0x05
