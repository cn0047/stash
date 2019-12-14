Makefile
-

````sh
make -f ed/sh/sh.makefile/example/one/Makefile etrp
make -f ed/sh/sh.makefile/example/one/Makefile var
````

Makefiles contain five kinds of things:

* explicit rules,
* implicit rules,
* variable definitions,
* directives, and comments.

The `@` before echo means - don't print this command to stdout before running it.

````sh
ifndef GOPATH
$(error ERROR!! GOPATH must be declared.)
endif

ifeq ($(CC),gcc)
endif
````
