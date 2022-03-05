# l1
First attempt at a simple Lisp in Go

WIP toy Lisp for fun projects.  Currently lexing works, parsing is next:

        12:37:03 l1 51.8F     ≡ * ☐ ~  >  go test
        PASS
        ok  	github.com/eigenhombre/l1	0.148s
        12:38:18 l1 51.8F     ≡ * ☐ ~  >  go build
        12:40:04 l1 52.7F     ≡ * ☐ ~  >  ./l1
        > (QUOTE (LAMBDA (X) (+ X X)))
        LPAREN('(')
        ATOM('QUOTE')
        LPAREN('(')
        ATOM('LAMBDA')
        LPAREN('(')
        ATOM('X')
        RPAREN(')')
        LPAREN('(')
        ATOM('+')
        ATOM('X')
        ATOM('X')
        RPAREN(')')
        RPAREN(')')
        RPAREN(')')

Take a look at the lexing tests and the rest of the code for a sense of what's there so far.
