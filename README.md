# l1

<img src="/l1.jpg" width="400">

![build](https://github.com/eigenhombre/l1/actions/workflows/build.yml/badge.svg)

First attempt at a simple Lisp in Go.

# Implemented
- Lexing and parsing
- Numbers as `big.Int`
- Eval
  - `t` as True, `()` as `nil`
  - Atoms bind to values in an environment
  - `quote`
  - Arithmetical operators `+-*/`
- (Some) simple error handling

# Examples

    $ l1
    > t
    t
    > ()
    ()
    > (quote foo)
    foo
    > (quote (the (ten (laws (of (greenspun))))))
    (the (ten (laws (of (greenspun)))))
    > (cdr (quote (is not common lisp)))
    (not common lisp)
    > (car (quote (is not common lisp)))
    is
    > 1
    1
    > -5
    -5
    > (* 12349807213490872130987 12349807213490872130987)
    152517738210391179737088822267441718485594169
    > (+)
    0
    > (+ 1 1 2 3)
    7
    > (+ 1 1)
    2
    > (eq (quote foo) (quote foo))
    t
    > (eq (quote foo) (quote bar))
    ()
    > (eq (quote foo) (quote (foo bar)))
    ()
    > (atom (quote foo))
    t
    > (atom (quote (foo bar)))
    ()
    > (
    unbalanced parens
    > )
    unexpected right paren
    > ^D
    $

These were copied directly from the unit test output; `eval_test.go` has more examples.

# Usage

Check out this repo and `cd` to it. Then,

- `go test` and maybe 
- `go build` followed by
- `go install`; then
- `l1`

A `Makefile` exists for convenience, and a `Dockerfile` for CI builds.

# Planned Features

- [Lisp 1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace);
- Sufficient power to [implement itself](http://www.paulgraham.com/rootsoflisp.html);
- Implement math as bignums from the get-go;
- Curses-based terminal control for text games, command line utilities, ...;

# Goals

- Improve my Go skills;
- Build a small, fast-loading Lisp that I can extend how I like;
- Learn more about [Lisp as a model for computation](http://www.paulgraham.com/rootsoflisp.html).

