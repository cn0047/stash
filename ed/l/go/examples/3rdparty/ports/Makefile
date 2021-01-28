export GOPATH=${CURDIR}

init:
	docker run -it --rm -v ${CURDIR}:/prj -w /prj -e GOPATH="/prj" cn007b/go \
	sh -c "go get -v -t -d ./..."

test:
	docker run -it --rm -v ${CURDIR}:/prj -w /prj -e GOPATH="/prj" cn007b/go \
	sh -c "go test -v ports/ddd/infrastructure/file_reader"

run:
	docker-compose up
