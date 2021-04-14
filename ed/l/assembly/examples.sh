# assembly

f=ed/l/assembly/examples/whatever/hw.asm

nasm -f macho64 -o x.o $f
ld -macosx_version_min 10.7.0 -no_pie -o x x.o && rm x.o
./x
