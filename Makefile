.PHONY: test clean deps lint all
.PHONY: verbose doc l1-tests release

PROG=l1
VERSION=`git describe --tags --abbrev=0`
LDFLAGS=-ldflags "-X main.version=${VERSION}"

all: deps test ${PROG} l1-tests lint doc

deps:
	go get .

${PROG}: *.go
	echo ${LDFLAGS}
	go build ${LDFLAGS} .

test:
	go test

l1-tests: ${PROG}
	./l1 tests.l1
	./l1 examples/fact.l1
	./l1 examples/fails.l1 && exit 1 || true
	./l1 examples/primes.l1
	./l1 examples/tco.l1

lint:
	golint -set_exit_status .
	staticcheck .

clean:
	rm -f ${PROG}

install: ${PROG}
	go install .

verbose: all
    # The tests are fast!  Just do it again, verbosely:
	go test -v

docker:
	docker build -t ${PROG} .

doc:
	python updatereadme.py

release:
	./bumpver
	make clean
	make
