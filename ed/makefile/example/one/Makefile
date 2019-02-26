TOP := "TOP"
ROOT = "ROOT"
TOP_ROOT_PATH=$(TOP):$(ROOT)/path

export GOPATH=${CURDIR}

etrp:
	echo $(TOP_ROOT_PATH)

var:
	x=$(shell echo 200)
	$(info VAR is $(x))

loop:
	@echo -e "1 \n2 \n3 \n4 \n5" | while read t; do echo $$t; done
