# The `l1` Language

Expressions in `l1` are atoms, lists, numbers, or functions.

Atoms have a name, such as `x`, `foo`, or `Eisenhower!`, and can be
"bound to an expression in the current environment," meaning,
essentially, given a value. For example,

    > (def a 1)
    > a
    1
    >

Lists are collections of zero or more expressions.  Examples:

    (getting atomic)
    (10 9 8 7 6 5 4 3 2 1)
    (+ 2 2)
    ()

Note: `(+ 2 2)` is a list that evaluates to the number `4`.  To
prevent evaluation of a list, prepend a quote character:

    > '(+ 2 2)
    (+ 2 2)
    > (+ 2 2)
    4

`()` is alternatively called the empty list, or nil.

Numbers are integer values and can be of arbitrary magnitude:

    0
    999
    7891349058731409803589073418970341089734958701432789

Numbers evaluate to themselves:

    > 0
    0
    > 7891349058731409803589073418970341089734958701432789
    7891349058731409803589073418970341089734958701432789

Numbers and atoms can be turned into lists:

    > (split atomic)
    (a t o m i c)
    > (split 1234)
    (1 2 3 4)

And lists can be turned into atoms:

    > (fuse '(getting atomic))
    gettingatomic
    > (fuse '(10 9 8 7 6 5 4 3 2 1))
    10987654321

# Boolean Logic

In `l1`, `()` is the only "falsey" value; everything else is "truthy".
Falsey and truthy are important when evaluating conditional statements
such as `if`, `when`, or `cond`.  The default truthy value is `t`.
`t` and `()` evaluate to themselves.

The `and`, `or` and `not` operators work like they do in most other
languages:

    > (and t t)
    t
    > (or () ())
    ()
    > (or () 135987)
    135987
    > (and () (launch missiles))
    ()
    > (not t)
    ()
    > (not ())
    t
    > (not 13987)
    ()
    > (if ()
        (launch missiles)
        555)
    555

# Special characters

Unlike many modern languages, `l1` doesn't have strings.  Instead,
atoms and lists are used where strings normally would be:

    > (printl '(Hello, world!))
    Hello, world!
    ()

(In this example, the `Hello, world!` is output to the terminal, and
then the return value of `printl`, namely `()`.)

Some characters, like `!`, need special handling, since they are "unreadable":

    > (printl '(!Hello, world!))
    unexpected lexeme 'unexpected character '!' in input'
    > (printl `(~(fuse `(~BANG Hello,)) world!))
    !Hello, world!
    ()

This is admittedly awkward, but rare in practice for the kinds of
programs `l1` was designed for.  `BANG` is one of a small set of atoms
helpful for this sort of construction:

    SPACE
    NEWLINE
    TAB
    BANG
    QMARK
    PERIOD
    COMMA

These all evaluate to atoms whose names are the unreadable characters.

# Functions

Functions come in two flavors: temporary functions, called "lambda"
functions for historical reasons, and functions which are defined and
kept around in the environment for later use.  For example,

    > (defn plus2 (x) (+ x 2))
    > (plus2 3)
    5
    > ((lambda (x) (* 5 x)) 3)
    15

Functions can take a fixed number of arguments plus an extra "rest"
argument, separated from the fixed arguments with a "."; the rest
argument is then bound to a list of all remaining arguments:

    > (defn multiply-then-sum (multiplier . xs)
        (* multiplier (apply + xs)))
    ()
    > (multiply-then-sum 5 1)
    5
    > (multiply-then-sum 5 1 2 3)
    30

A functions that has a rest argument but no fixed arguments is
specified using the empty list as its fixed argument:

    > (defn say-hello (() . friends)
        (list* 'hello friends))
    > (say-hello 'John 'Jerry 'Eden)
    (hello John Jerry Eden)


