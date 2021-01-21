Assembly (Assembler)
-
Since 1949.

````sh
*.s # source code files written in assembly
````

General purpose registers:
* 16 bits:
  * AL, BL, CL, DL; R8B
  * AH, BH, CH, DH
* 32 bits:
  *  AX,  BX,  CX,  DX;  SI,  DI,  BP,  SP; R8W
  * EAX, EBX, ECX, EDX; ESI, EDI, EBP, ESP; R8D
* 64 bits:
  * RAX, RBX, RCX, RDX; RSI, RDI, RBP, RSP; R8-R15

````
EIP/RIP - instruction pointer.
ESP/RSP - stack pointer.
EBP/RBP - base pointer.
ESI/RSI - source index.
EDI/RDI - destination index.

CF - carry flag
ZF - zero flag
SF - sign flag
TF - trap flag
DF - direction flag
OF - overflow flag

JE/JZ   - jump equal | zero
JNE/JNZ - jump not equal | not zero
JG/JNLE - jump greater | not less/equal
JGE/JNL - jump greater/equal | not less
JL/JNGE - jump less | not qreater/equal
JLE/JNG - jump less/equal | not greater

bit         - 1/0
byte        - 8 bits
word        - 2 bytes
double-word - 4 bytes
quad-word   - 8 bytes
````

<br>ISA - Instruction Set Architecture.
<br>Big-Endian - the most significant byte first.
<br>Little-Endian - the last significant byte first.

````
ADD DEST, SOURCE
SUB DEST, SOURCE
````
