.PHONY: test clean deps lint all verbose doc

PROG=l1

all: deps test ${PROG} lint doc

deps:
	go get .

${PROG}: *.go
	go build .

test:
	go test

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
