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

**.PHONY** means not associated with files
(`.PHONY: install` won't check whether the install file exists).

````sh
.PHONY: clean
clean:
  rm -rf *.o

VAR_F ?= "foo" # set VAR_F if it not set.

ifndef GOPATH
$(error ERROR!! GOPATH must be declared.)
endif

ifeq ($(CC),gcc)
endif
````
