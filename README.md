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
        > 1
        1
        > 31489071430987532109487513094875031984750983147
        31489071430987532109487513094875031984750983147
        > (quote (the (ten (laws (of (greenspun))))))
        (the (ten (laws (of (greenspun)))))
        > (car (quote (the (ten (laws (of (greenspun)))))))
        the
        > (cdr (quote (is not common lisp)))
        (not common lisp)
        > t
        t
        > ()
        ()
        > (* (+ 1 2 3) (/ 4000 1000) 2139872138907)
        51356931333768
        > (
        unbalanced parens
        > )
        unexpected right paren
        > ^D
        $

Take a look at the `eval_test.go` for a better sense of what's implemented so far.

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

