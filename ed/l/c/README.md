c
-

````sh
# osx
# -Wall - warnings all
gcc -o x -Wall ed/l/c/examples/whatever/hw.c && ./x

# osx
cc -v

*.c          # source code
*.obj or *.o # object code
````

<br>preprocessor - optimize source code and pass it to compiler.
<br>compiler - .
<br>gcc - GNU Compiler Collection.
<br>linker - link all `*.obj` files.

Data types:
* int (short, long)
* float
* double
* char

**string** `char str[] = "my string";`

````c
unsigned int x = 1u;
short int y = 2;
long int z = 3;
byte b = 0x12;

if (exp)
if (exp) st else st
if (exp) st else if (exp) st
switch (exp) {
  case cond: st break;
  default: st
}

while (exp) {}
do {} while(exp)
for (ifexp; condexp; loopexp) {}
for (;;) {}
````
