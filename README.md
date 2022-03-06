# l1
First attempt at a simple Lisp in Go

WIP toy Lisp for fun projects.  Currently lexing and parsing works, eval is next:

        11:57:20 l1 39.3F     ≡ * ☐ ~ (master) >  ./l1
        > (QUOTE (LAMBDA (X) (+ X X)))

        LEXEMES_____________
        LP('(')
        ATOM('QUOTE')
        LP('(')
        ATOM('LAMBDA')
        LP('(')
        ATOM('X')
        RP(')')
        LP('(')
        ATOM('+')
        ATOM('X')
        ATOM('X')
        RP(')')
        RP(')')
        RP(')')

        PARSED ITEMS________
        [(QUOTE (LAMBDA (X) (+ X X)))]
        > ^C
        11:57:31 l1 39.3F     ≡ * ☐ ~ (master) >

Take a look at the lexing tests and the rest of the code for a sense of what's there so far.
