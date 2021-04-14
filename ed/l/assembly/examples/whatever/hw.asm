; This is first Hello World program.

global start

section .data
  message: db "Hello, World", 10

section .text
start:
  mov rax, 0x02000004
  mov rdi, 1
  mov rsi, message
  mov rdx, 13
  syscall

  mov rax, 0x02000001
  xor rdi, rdi
  syscall
