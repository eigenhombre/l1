.PHONY: test clean deps lint all verbose

PROG=l1

all: deps test ${PROG} lint

deps:
	go get .

${PROG}: *.go
	go build .

test:
	go test

lint:
	golint -set_exit_status .
# Broken in Go 1.18: (l1 issue #24):
#	staticcheck .

clean:
	rm ${PROG}

install: ${PROG}
	go install .

verbose: all
    # The tests are fast!  Just do it again, verbosely:
	go test -v
