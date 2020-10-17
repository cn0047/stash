c
-

[c debugger](https://web.eecs.umich.edu/~sugih/pointers/summary.html)

````sh
# -Wall - warnings all
gcc -o x -Wall $f

clang -v
clang -o x $f && ./x

# osx
cc -v
cc -o x $f && ./x

*.c          # source code
*.obj or *.o # object code
````

<br>preprocessor - optimize source code and pass it to compiler.
<br>compiler - .
<br>cc - C Compiler.
<br>gcc - GNU Compiler Collection.
<br>linker - link all `*.obj` files.

#### Data types

* int (unsigned, short, long)
* float
* double
* char

* enum
* typedef

* array
* struct
* pointer
* union

````c
const int x = 1;

printf("%s\n", str);
printf("%f\n", float);
printf("%p\n", &pointer);

unsigned int x = 1u;
short int y = 2;
long int z = 3;
byte b = 0x12;

if (exp) st else if (exp) st else st
switch (exp) {
  case cond: st break;
  default: st
}

while (exp) {}
do {} while(exp)
for (int i=0; i<10; i++) {}

malloc()  // Allocate a block of memory
calloc()  // Allocate a memory block and fill it with null bytes
realloc() // Resize an allocated memory block
free()    // Release a memory block
````

**string** `char str[] = "my string";`

**union** special data type that allows store different data types in the same memory location.
