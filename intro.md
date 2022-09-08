# The `l1` Language

[Jump to API](#api-index) / list of operators

## The REPL

Examples in these docs are shown as if typed at the `l1` "REPL"
(read-eval-print-loop).  To leave the REPL, type Control-D or
`(exit)`:

    $ l1
    > (exit)
    $

If you want arrow keys for next/previous command or to move forward or
backward on the line, history, etc., wrap `l1` with the `rlwrap`
command (installed using your favorite package manager), e.g.:

    $ rlwrap l1
    >

## Expressions

Expressions in `l1` are atoms, lists, numbers, or functions:

### Atoms

Atoms have a name, such as `x`, `foo`, or `Eisenhower!`, and can be
"bound to a value in the current environment," meaning,
essentially, assigned a value. For example, we can bind a value to `a`
and retrieve it:

    > (def a 1)
    > a
    1
    >

### Lists

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

As in most Lisps, lists are actually implemented
"under the hood" as "cons pairs" ([Wikipedia
page](https://en.wikipedia.org/wiki/Cons#Ordered_pairs)).  The list

    (1 2 3)

actually represented internally as

    (1 . (2 . (3 . ())))

where

    (a . b)

is the same as

    (cons a b)

In practice, the dot notation is uncommon in `l1` programs, except
when used to represent rest arguments, described below.

### Numbers

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

Atom names can be arbitrarily long (they are simply Go strings under
the hood).  When `l1` parses your code, it will interpret any UTF-8-encoded unicode characters
but the following as the start of an atom:

    0123456789+-. \t\n\r()~@#;`'

After the first character, anything is allowed except spaces or

    \t\n\r()~@#;`'

Deviations from these constraints need special handling.  For example:

    > (printl '(@Hello, world!))
    (...
     (unexpected lexeme 'unexpected character '@' in input' on line 1))

A workaround is to use syntax quote and unquote, described below, to
dynamically create a new atom name using the `BANG` alias for `!`:

    > (printl `(~(fuse `(~ATSIGN Hello,)) world!))
    @Hello, world!
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
functions for historical reasons, and functions which are defined with
a name and kept around in the environment for later use.  For example,

    > (defn plus2 (x) (+ x 2))
    > (plus2 3)
    5
    > ((lambda (x) (* 5 x)) 3)
    15

Since function names are atoms, their names follow the same rules for
atoms given above.  The following function definitions are all equally valid:

    > (defn increase! (x) (* x 1000))
    ()
    > (defn 增加! (x) (* x 1000))
    ()
    > (defn մեծացնել! (x) (* x 1000))
    ()

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

In addition to the functions described above, some `l1` functions are
"built in" (implemented in Go as part of the language core).  Examples
include `car`, `cdr`, `cons`, etc.  The API Docs below specify whether
a function is built-in or not.

One special family of predefined functions not shown in the API docs
(because they are effectively infinite in number) is extensions of
`car` and `cdr`:

    > (car '(1 2 3))
    1
    > (cadr '(1 2 3))
    2
    > (caddr '(1 2 3))
    3
    > (caar '((one fish) (two fish)))
    one
    > (caadr '((one fish) (two fish)))
    two
    > (cadar '((one fish) (two fish)))
    fish

`(cadr x)` should be read as `(car (cdr x))`, and so on.

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

## Flow of Control

In addition to the basic conditional statements `cond`, `if`,
`if-not`, `when`, and `when-not`, flow of control is generally
implemented via recursion, as it is in Scheme, and inspection of its
core library
[`l1.l1`](https://github.com/eigenhombre/l1/blob/master/l1.l1) will
show several examples of recursive functions being used as the primary
recurrence method. A few other control flow methods are also
available: [`while`](#while), which loops so long as a condition is
true; [`dotimes`](#dotimes), which executes a body of statements a
given number of times; [`foreach`](#foreach), which executes a body of
statements for each element in a loop; and [`loop`](#loop), which
loops forever.  Macros can be used to create new control abstractions;
however, possibilities are somewhat restricted compared to some
languages, due to the inability of `l1` to handle branches or "goto"
statements.

## Assertions and Error Handling

### `is`

There is currently one kind of assertion expression in `l1`, namely
the `is` macro:

    > (is (= 4 (+ 2 2)))
    > (is ())
    ERROR:
    ((assertion failed: ()))
    >

If the argument to `is` is falsey (`()`), then an error is printed
and the program exits (if running a program) or the REPL
returns to the top level prompt.

If `is` is checking an equality of two items which fails, the macro is smart
enough to print out a more detailed error report showing the two
expressions and their values:

    > (is (= 5 (+ 1 1)))
    ERROR:
    ((expression 5 ==> 5 is not equal to expression (+ 1 1) ==> 2))
    >

### `error`

If desired, an error can be caused deliberately with the `error` function:

    > (defn checking-len (x)
        (if (list? x)
         (len x)
         (error '(argument must be a list))))
    ()
    > (checking-len 3)
    ERROR:
    ((argument must be a list))
    >

### `errors`

In some situations, such as during automated tests, it may be
desirable to ensure that a specific error is raised.  The `errors`
special form takes a list argument and a body, then checks to ensure
that the body raises an error which contains the supplied list:

    > (errors '(division by zero) (/ 1 0))
    ()
    > (errors '(division) (/ 1 0))
    ()
    > (errors '(rocket crashed) (/ 1 0))
    ERROR:
    ((error 'rocket crashed' not found in '((builtin function /) (division by zero))'))
    > (errors '(division by zero) (* 1 0))
    ERROR:
    ((error not found in ((quote (division by zero)) (* 1 0))))
    >

### `try ... catch`

Most contemporary languages will print a stacktrace when an error
occurs.  `l1` stacktraces are somewhat rudimentary: in keeping with
the rest of the language, they are simply lists.  To capture an error
occurring within a body of code, wrap the body in a `try` statement
and add a `catch` clause, as follows:

    > (try
        (printl '(got here))
        (+ 3
           (/ 1 0))
        (printl '(did not get here))
      (catch e
        (cons '(oh boy, another error)
              e)))
    got here
    ((oh boy, another error) (builtin function /) (division by zero))
    >

The exception `e` is a list of lists to which items are added (in the
front) as the error returns up the call chain.  As an ordinary list,
it can be manipulated like any other, as shown above using `cons`.

An important caveat is that, since tail recursion is optimized away,
many "stack frames" (or their equivalent) are optimized away - there
is no way to track the entire history in detail without losing the
space-saving power of the optimization.  Nevertheless, the generated
exception can be helpful for troubleshooting.

There is, currently, no equivalent of the `finally` clause one sees in
Java or Clojure.

### `swallow`

Rarely, one may wish to swallow any errors and continue execution.
`swallow` will execute all the statements in its body and return `t`
if and only if any of them causes an (uncaught) error:

    > (swallow (/ 1 0))
    t
    > (swallow (+ 1 1))
    ()
    >

`swallow` is used mainly in the fuzzing tests for `l1` (see the
examples directory).

## Subprocesses

The `shell` function executes a subprocess command, which should be a
list of atoms and numbers, and returns the result in the following
form:

    ((... stdout lines...)
     (... stderr lines...)
     ...exit code..)

Examples (output reformatted for clarity):

    > (shell '(pwd))
    (((/Users/jacobsen/Programming/go/l1))
     (())
     0)
    > (shell '(ls))
    (((Dockerfile) (LICENSE) (Makefile) (README.md) (api.md)
      (atom.go) (builtin.go) (builtin_test.go) (bumpver)
      (cons.go) (core.go) (doc.go) (env.go) (env_test.go)
      (error.go) (error_test.go) (eval_test.go) (example.l1)
      (examples) (examples.txt) (go.mod) (go.sum) (intro.md)
      (l1) (l1.html) (l1.jpg) (l1.l1) (l1.md) (lambda.go)
      (lex.go) (lex_test.go) (lisp.go) (main.go) (math.go)
      (math_test.go) (parse.go) (parse_test.go) (sexpr_test.go)
      (shell.go) (state.go) (term.go) (tests.l1) (updatereadme.py)
      (util.go) (version.go))
     (())
     0)
    > (shell '(ls -al l1.l1))
    (((-rw-r--r-- 1 jacobsen staff 16636 Sep 3 11:28 l1.l1)) (()) 0)
    > (shell '(ls /watermelon))
    ((()) ((ls: /watermelon: No such file or directory)) 1)

## Macros

For those familiar with macros (I recommend Paul
Graham's *On Lisp* for those who are not), `l1` macros are
[non-hygienic](https://en.wikipedia.org/wiki/Hygienic_macro) by
default.  Gensyms, unique atom names useful for writing safe macros,
are available via the `gensym` built-in function:

    > (gensym)
    <gensym-0>
    > (gensym 'foo)
    <gensym-foo-1>

The traditional `syntax-quote`, `unquote`, and `splicing-unquote` are available, and
have sugared equivalents:

    (let ((name 'larry)
          (names '(moe curly)))
      (syntax-quote (hello, (unquote name) as well as (splicing-unquote names))))
    ;;=>
    (hello, larry as well as moe curly)

is the same as

    (let ((name 'larry)
          (names '(moe curly)))
      `(hello, ~name as well as ~@names))

In addition to the quote (`'`), syntax-quote (`` ` ``), unquote (`~`),
and splicing unquote (`~@`) shortcuts, the shortcut `#_` is available, which is equivalent to `(comment ...)`, e.g.,

    #_(this is a commented form)

is equivalent to

    (comment (this is a commented form))

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

## Running l1 Programs as Command Line Scripts

Programs can be run by giving the program name as an argument to `l1`:

    l1 hello.l1

However, if you add `#!/usr/bin/env l1` at the beginning of an `l1` file:

    #!/usr/bin/env l1
    ;; hello.l1
    (printl '(hello world))

and set the execute bit on the file permissions:

    chmod +x hello.l1

then you can run `hello.l1` "by itself," without explicitly invoking `l1`:

    $ ./hello.l1
    hello world
    $

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
