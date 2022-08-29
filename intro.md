# The `l1` Language

[Jump to API](#api-index) / list of operators

## Expressions

Expressions in `l1` are atoms, lists, numbers, or functions.

Atoms have a name, such as `x`, `foo`, or `Eisenhower!`, and can be
"bound to a value in the current environment," meaning,
essentially, assigned a value. For example, we can bind a value to `a`
and retrieve it:

    > (def a 1)
    > a
    1
    >

Lists are collections of zero or more expressions.  Examples:

    (getting atomic)
    (10 9 8 7 6 5 4 3 2 1)
    (+ 2 2)
    ()

In general, lists represent operations whose name is the first element
and whose arguments are the remaining elements.  For example, `(+ 2
2)` is a list that evaluates to the number `4`.  To prevent evaluation
of a list, prepend a quote character:

    > '(+ 2 2)
    (+ 2 2)
    > (+ 2 2)
    4

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

    > (split 'atomic)
    (a t o m i c)
    > (split 1234)
    (1 2 3 4)

And lists can be turned into atoms or numbers:

    > (fuse '(getting atomic))
    gettingatomic
    > (fuse '(10 9 8 7 6 5 4 3 2 1))
    10987654321

## Boolean Logic

In `l1`, the empty list `()` is the only "falsey" value; everything
else is "truthy".  Falsey and truthy are important when evaluating
conditional statements such as `if`, `when`, or `cond`.  The default
truthy value is `t`.  `t` and `()` evaluate to themselves.

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

## Special Characters

Unlike many modern languages, `l1` doesn't have strings.  Instead,
atoms and lists are used where strings normally would be:

    > (printl '(Hello, world!))
    Hello, world!
    ()

(In this example, the `Hello, world!` is output to the terminal, and
then the return value of `printl`, namely `()`.)

Atom names can be arbitrarily long (they are Go strings under the hood)
and can start with the characters

    abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+*/-=_<>?

After the first character,

    0123456789!$^.,

are also allowed.  Deviations from these constraints need special
handling.  For example:

    > (printl '(!Hello, world!))
    unexpected lexeme 'unexpected character '!' in input'

A workaround is to use syntax quote and unquote, described below, to
dynamically create a new atom name using the `BANG` alias for `!`:

    > (printl `(~(fuse `(~BANG Hello,)) world!))
    !Hello, world!
    ()

This is admittedly awkward, but rare in practice for the kinds of
programs `l1` was designed for.  `BANG` is one of a small set of atoms
helpful for this sort of construction:

    BANG
    COLON
    COMMA
    NEWLINE
    PERIOD
    QMARK
    SPACE
    TAB

These all evaluate to atoms whose names are the unreadable characters,
some of which may be helpful for text games and other diversions:

    > (dotimes 10
        (println
         (fuse
          (repeatedly 10
                      (lambda ()
                        (randchoice (list COMMA
                                          COLON
                                          PERIOD
                                          BANG
                                          QMARK)))))))

    .!!!.???..
    ,??::?!,.?
    ?,?!?..:!!
    ,:.,?.:!!!
    !!:?!::.,?
    ,:!!!:,!!:
    ,???:?!:!?
    .,!!?,!:!?
    !:,!!!.:!:
    ??.,,:.:..
    ()
    >

## Functions

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

A function that has a rest argument but no fixed arguments is
specified using the empty list as its fixed argument:

    > (defn say-hello (() . friends)
        (list* 'hello friends))
    > (say-hello 'John 'Jerry 'Eden)
    (hello John Jerry Eden)

Functions may invoke themselves recursively:

    > (defn sum-nums (l)
        (if-not l
          0
          (+ (car l) (sum-nums (cdr l)))))
    ()
    > (sum-nums '(0 1 2 3 4 5 6 7 8 9))
    45

The above function performs an addition after it invokes itself.  A
function which invokes itself *immediately before returning*, without
doing any more work, is called "tail recursive."  Such functions are
turned into iterations automatically by the interpreter ("tail call
optimization").  The above function can be rewritten into a
tail-recursive version:

    > (defn sum-nums-accumulator (l acc)
        (if-not l
          acc
          (sum-nums-accumulator (cdr l) (+ acc (car l)))))
    ()
    > (sum-nums-accumulator '(0 1 2 3 4 5 6 7 8 9) 0)
    45

Lambda functions can invoke themselves if given a name directly before
the parameters are declared.  We can rewrite the above function to
hide the `acc` argument from the user:

    > (defn sum-nums (l)
        (let ((inner (lambda inner (l acc)
                       (if-not l
                         acc
                         (inner (cdr l) (+ acc (car l)))))))
          (inner l 0)))
    ()
    > (sum-nums '(0 1 2 3 4 5 6 7 8 9))
    45

In this version, `inner` is tail-recursive, and `sum-nums` is now as
convenient to use as our first, non-tail-recursive version was.

## Macros

For those familiar with macros (if you're not I recommend Paul
Graham's *On Lisp*), `l1` macros are
[non-hygienic](https://en.wikipedia.org/wiki/Hygienic_macro) by
default and there is no built-in gensym utility
([yet](https://github.com/eigenhombre/l1/issues/63)).  Collisions can
be made arbitrarily unlikely by constructing random atom names:

    (defn gensym ()
      (fuse (concat (randalpha 5) (randigits 5))))

The traditional `syntax-quote`, `unquote`, and `splicing-unquote` are available, and
have sugared equivalents:

    (let ((name 'larry)
          (names '(moe curly)))
      (syntax-quote (hello, (unquote name) as well as (splicing-unquote names))))

is the same as

    (let ((name 'larry)
          (names '(moe curly)))
      `(hello, ~name as well as ~@names))

In addition to the quote (`'`), syntax-quote (`` ` ``), unquote (`~`),
and splicing unquote (`~@`) shortcuts, the shortcut `#_` is available, which is equivalent to `(comment ...)`, e.g.,

    #_(this is a commented form)

is equivalent to

    (comment (this is a commented form))

## Running l1 Programs as Command Line Scripts

Programs can be run by giving the program name as an argument to `l1`:

    l1 hello.l1

However, if you add `#!/usr/bin/env l1` at the beginning of an `l1` file:

    #!/usr/bin/env l1
    ;; hello.l1
    (printl '(hello world))

and set the execute bit on the file permissions:

    chmod +x hello.l1

then you can run hello.l1 "by itself," without explicitly invoking `l1`:

    $ ./hello.l1
    hello world
    $

## Text User Interfaces

`l1` has a few built-in functions for creating simple text UIs:

- `screen-clear`: Clear the screen
- `screen-get-key`: Get a keystroke
- `screen-write`: Write a list, without parentheses, to an `x` and `y` position on the screen.
- `with-screen` (macro): Enter/exit "screen" (UI) mode

The `screen-...` functions must occur within a `with-screen`
expression.  [An example
program](https://github.com/eigenhombre/l1/blob/master/examples/screen-test.l1)
shows these functions in action.

## Emacs Integration

If you are using Emacs, you can set it up to work with `l1` as an "inferior
lisp" process as described in [the Emacs manual](https://www.gnu.org/software/emacs/manual/html_node/emacs/External-Lisp.html).
I currently derive a new major mode from the base `lisp-mode` and bind a few
keys for convenience as follows:

    (define-derived-mode l1-mode
      lisp-mode "L1 Mode"
      "Major mode for L1 Lisp code"
      (setq inferior-lisp-program (executable-find "l1")
      (paredit-mode 1)
      (put 'test 'lisp-indent-function 1)
      (put 'testing 'lisp-indent-function 1)
      (put 'errors 'lisp-indent-function 1)
      (put 'if 'lisp-indent-function 1)
      (put 'if-not 'lisp-indent-function 1)
      (put 'foreach 'lisp-indent-function 2)
      (put 'when-not 'lisp-indent-function 1)
      (define-key l1-mode-map (kbd "s-i") 'lisp-eval-last-sexp)
      (define-key l1-mode-map (kbd "s-I") 'lisp-eval-form-and-next)
      (define-key l1-mode-map (kbd "C-o j") 'run-lisp))

    (add-to-list 'auto-mode-alist '("\\.l1" . l1-mode))

If `l1` has been installed on your path, `M-x run-lisp` or using the appropriate
keybinding should be enough to start a REPL within Emacs and start sending
expressions to it.
