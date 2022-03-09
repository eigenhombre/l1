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
- Simple error handling

# Examples

        $ l1
        > 1
        1
        > 31489071430987532109487513094875031984750983147
        31489071430987532109487513094875031984750983147
        > (quote (lists exist but are not evaluated yet))
        (lists exist but are not evaluated yet)
        > t
        t
        > ()
        ()
        > (
        unbalanced parens
        > )
        unexpected right paren
        > ^D
        $

Take a look at the tests for a better sense of what's there so far.

# Usage

Check out this repo and `cd` to it. Then,

- `go test` and maybe 
- `go build` followed by
- `go install`; then
- `l1`

A `Makefile` exists for convenience, and a `Dockerfile` for CI builds.

# L1 features (planned)

- [Lisp 1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace);
- Sufficient power to [implement itself](http://www.paulgraham.com/rootsoflisp.html);
- Implement math as bignums from the get-go;
- Curses-based terminal control for text games, command line utilities, ...;

# Goals

- Improve my Go skills;
- Build a small, fast-loading Lisp that I can extend how I like;
- Learn more about [Lisp as a model for computation](http://www.paulgraham.com/rootsoflisp.html).

