.PHONY: test clean deps lint all
.PHONY: verbose doc integration-test

PROG=l1

all: deps test ${PROG} integration-test lint doc

deps:
	go get .

${PROG}: *.go
	go build .

test:
	go test

integration-test: ${PROG}
	./l1 fact.l1
	./l1 fails.l1 && exit 1 || true

lint:
	golint -set_exit_status .
	staticcheck .

clean:
	rm ${PROG}

install: ${PROG}
	go install .

verbose: all
    # The tests are fast!  Just do it again, verbosely:
	go test -v

docker:
	docker build -t ${PROG} .

doc:
	python updatereadme.py
