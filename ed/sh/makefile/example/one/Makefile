# @example:
# make -f ed/sh/makefile/example/one/Makefile etrp
# make -f ed/sh/makefile/example/one/Makefile etrp2
# make -f ed/sh/makefile/example/one/Makefile etrp_all
# make -f ed/sh/makefile/example/one/Makefile var

TOP := "TOP"
ROOT = "ROOT"
TOP_ROOT_PATH=$(TOP):$(ROOT)/path

export GOPATH=${CURDIR}

etrp:
	echo 'TOP_ROOT_PATH='$(TOP_ROOT_PATH)

etrp2:
	@echo "\nTOP_ROOT_PATH=$(TOP_ROOT_PATH)\n"

etrp_all: etrp etrp2

print_f:
	@echo $(@F) # will print: print_f

var:
	x=$(shell echo 200)
	$(info VAR is $(x))

.PHONY: var1
var1:
	x=$(shell echo 200)
	$(info VAR is $(x))

var2:
	@echo 501 > /tmp/make.test
	$(eval c := $(shell cat /tmp/make.test))
	@echo $(c) > /tmp/make.test.res
	@cat /tmp/make.test.res

loop:
	@echo -e "1 \n2 \n3 \n4 \n5" | while read t; do echo $$t; done

.PHONY: cmd1
cmd1: x=204
cmd1: cmd

.PHONY: cmd
cmd:
	@echo $(x)
