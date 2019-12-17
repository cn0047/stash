TOP := "TOP"
ROOT = "ROOT"
TOP_ROOT_PATH=$(TOP):$(ROOT)/path

export GOPATH=${CURDIR}

etrp:
	echo $(TOP_ROOT_PATH)

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
