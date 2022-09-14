.PHONY: test clean deps lint all
.PHONY: verbose doc l1-tests release
.PHONY: tco-test slow fast run-examples

PROG=l1

all: deps fast slow

slow: tco-test

fast: test ${PROG} l1-tests run-examples lint doc

deps:
	go get ./src

${PROG}: src/*.go src/l1.l1
	go build -o l1 ./src

test:
	go test ./src

l1-tests: ${PROG}
	./l1 tests.l1
	./l1 -e "(println (+ 1 1))"
	./l1 -e "(load 'examples/fact.l1)"
	./l1 -e "(error '(goodbye, cruel world))" && exit 1 || echo "Got expected error"

run-examples: ${PROG}
	./l1 examples/fact.l1
	./l1 examples/fails.l1 && exit 1 || true
	./l1 examples/primes.l1
	./l1 examples/sentences.l1
	./l1 examples/galax.l1

tco-test: ${PROG}
	./l1 examples/tco.l1

lint:
	golint -set_exit_status ./src
	staticcheck ./src

clean:
	rm -f ${PROG}

install: ${PROG}
	go install .

verbose: all # The tests are fast!  Just do it again, verbosely:
	go test -v ./src

docker:
	docker build -t ${PROG} .

doc:
	./l1 -longdoc > api.md
	cat intro.md api.md > l1.md
	python updatereadme.py

release:
	./bumpver
	make clean
	make
