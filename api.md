# API Index
110 forms available:
[`*`](#*)
[`**`](#**)
[`+`](#+)
[`-`](#-)
[`/`](#/)
[`<`](#<)
[`<=`](#<=)
[`=`](#=)
[`>`](#>)
[`>=`](#>=)
[`and`](#and)
[`apply`](#apply)
[`atom?`](#atom?)
[`bang`](#bang)
[`body`](#body)
[`butlast`](#butlast)
[`caar`](#caar)
[`capitalize`](#capitalize)
[`car`](#car)
[`cdr`](#cdr)
[`colon`](#colon)
[`comma`](#comma)
[`comment`](#comment)
[`complement`](#complement)
[`concat`](#concat)
[`cond`](#cond)
[`cons`](#cons)
[`constantly`](#constantly)
[`dec`](#dec)
[`def`](#def)
[`defmacro`](#defmacro)
[`defn`](#defn)
[`doc`](#doc)
[`dotimes`](#dotimes)
[`downcase`](#downcase)
[`drop`](#drop)
[`error`](#error)
[`errors`](#errors)
[`even?`](#even?)
[`exclaimed`](#exclaimed)
[`filter`](#filter)
[`flatten`](#flatten)
[`foreach`](#foreach)
[`forms`](#forms)
[`fuse`](#fuse)
[`help`](#help)
[`identity`](#identity)
[`if`](#if)
[`if-not`](#if-not)
[`inc`](#inc)
[`is`](#is)
[`lambda`](#lambda)
[`last`](#last)
[`len`](#len)
[`let`](#let)
[`list`](#list)
[`list*`](#list*)
[`list?`](#list?)
[`loop`](#loop)
[`macroexpand-1`](#macroexpand-1)
[`map`](#map)
[`mapcat`](#mapcat)
[`neg?`](#neg?)
[`not`](#not)
[`nth`](#nth)
[`number?`](#number?)
[`odd?`](#odd?)
[`or`](#or)
[`period`](#period)
[`pos?`](#pos?)
[`print`](#print)
[`printl`](#printl)
[`println`](#println)
[`progn`](#progn)
[`punctuate`](#punctuate)
[`punctuate-atom`](#punctuate-atom)
[`quote`](#quote)
[`randalpha`](#randalpha)
[`randchoice`](#randchoice)
[`randigits`](#randigits)
[`randint`](#randint)
[`range`](#range)
[`readlist`](#readlist)
[`reduce`](#reduce)
[`rem`](#rem)
[`remove`](#remove)
[`repeat`](#repeat)
[`repeatedly`](#repeatedly)
[`reverse`](#reverse)
[`screen-clear`](#screen-clear)
[`screen-end`](#screen-end)
[`screen-get-key`](#screen-get-key)
[`screen-size`](#screen-size)
[`screen-start`](#screen-start)
[`screen-write`](#screen-write)
[`shuffle`](#shuffle)
[`sleep`](#sleep)
[`some`](#some)
[`split`](#split)
[`syntax-quote`](#syntax-quote)
[`take`](#take)
[`test`](#test)
[`tosentence`](#tosentence)
[`true?`](#true?)
[`upcase`](#upcase)
[`version`](#version)
[`when`](#when)
[`when-not`](#when-not)
[`with-screen`](#with-screen)
[`zero?`](#zero?)
# Operators


## `*`

Multiply 0 or more numbers

Type: native function

Arity: 0+

Args: `(() . xs)`


### Examples

```
> (* 1 2 3)
;;=>
6
> (*)
;;=>
1

```

-----------------------------------------------------
		

## `**`

Exponentiation operator

Type: function

Arity: 2 

Args: `(n m)`


### Examples

```
> (** 1 0)
;;=>
1
> (** 2 4)
;;=>
16
> (** 10 10)
;;=>
10000000000

```

-----------------------------------------------------
		

## `+`

Add 0 or more numbers

Type: native function

Arity: 0+

Args: `(() . xs)`


### Examples

```
> (+ 1 2 3)
;;=>
6
> (+)
;;=>
0

```

-----------------------------------------------------
		

## `-`

Subtract 0 or more numbers from the first argument

Type: native function

Arity: 1+

Args: `(x . xs)`


### Examples

```
> (- 1 1)
;;=>
0
> (- 5 2 1)
;;=>
2
> (- 99)
;;=>
-99

```

-----------------------------------------------------
		

## `/`

Divide the first argument by the rest

Type: native function

Arity: 2+

Args: `(numerator denominator1 . more)`


### Examples

```
> (/ 1 2)
;;=>
0
> (/ 12 2 3)
;;=>
2
> (/ 1 0)
;;=>
ERROR: division by zero

```

-----------------------------------------------------
		

## `<`

Return t if the arguments are in strictly increasing order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


### Examples

```
> (< 1 2)
;;=>
t
> (< 1 1)
;;=>
()
> (< 1)
;;=>
t
> (apply < (range 100))
;;=>
t

```

-----------------------------------------------------
		

## `<=`

Return t if the arguments are in increasing (or qual) order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


### Examples

```
> (<= 1 2)
;;=>
t
> (<= 1 1)
;;=>
t
> (<= 1)
;;=>
t

```

-----------------------------------------------------
		

## `=`

Return t if the arguments are equal, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


### Examples

```
> (= 1 1)
;;=>
t
> (= 1 2)
;;=>
()
> (apply = (repeat 10 t))
;;=>
t

```

-----------------------------------------------------
		

## `>`

Return t if the arguments are in strictly decreasing order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


### Examples

```
> (> 1 2)
;;=>
()
> (> 1 1)
;;=>
()
> (> 1)
;;=>
t

```

-----------------------------------------------------
		

## `>=`

Return t if the arguments are in decreasing (or equal) order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


### Examples

```
> (>= 1 2)
;;=>
()
> (>= 1 1)
;;=>
t

```

-----------------------------------------------------
		

## `and`

Boolean and

Type: special form

Arity: 0+

Args: `(() . xs)`


### Examples

```
(and)
;;=>
true
> (and t t)
;;=>
true
> (and t t ())
;;=>
()
> (and () (/ 1 0))
;;=>
()

```

-----------------------------------------------------
		

## `apply`

Apply a function to a list of arguments

Type: native function

Arity: 2 

Args: `(f args)`


### Examples

```
> (apply + (repeat 10 1))
;;=>
10
> (apply * (cdr (range 10)))
;;=>
362880

```

-----------------------------------------------------
		

## `atom?`

Return t if the argument is an atom, () otherwise

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (atom? 1)
;;=>
()
> (atom? (quote one))
;;=>
t

```

-----------------------------------------------------
		

## `bang`

Add an exclamation point at end of atom

Type: function

Arity: 1 

Args: `(a)`


### Examples

```
> (bang (quote Bang))
;;=>
Bang!

```

-----------------------------------------------------
		

## `body`

Return the body of a lambda function

Type: native function

Arity: 1 

Args: `(f)`


### Examples

```
> (body (lambda (x) (+ x 1)))
;;=>
((+ x 1))

```

-----------------------------------------------------
		

## `butlast`

Return everything but the last element

Type: function

Arity: 1 

Args: `(l)`


### Examples

```
> (butlast ())
;;=>
()
> (butlast (range 3))
;;=>
(0 1)

```

-----------------------------------------------------
		

## `caar`

First element of the first element of a list of lists

Type: function

Arity: 1 

Args: `(l)`


### Examples

```
> (caar ())
;;=>
()
> (caar (quote (())))
;;=>
()
> (caar (quote ((one two) (three four))))
;;=>
one

```

-----------------------------------------------------
		

## `capitalize`

Return the atom argument, capitalized

Type: function

Arity: 1 

Args: `(a)`


### Examples

```
> (capitalize (quote hello))
;;=>
Hello

```

-----------------------------------------------------
		

## `car`

Return the first element of a list

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (car (quote (one two)))
;;=>
one
> (car ())
;;=>
()

```

-----------------------------------------------------
		

## `cdr`

Return a list with the first element removed

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (cdr (quote (one two)))
;;=>
(two)
> (cdr ())
;;=>
()

```

-----------------------------------------------------
		

## `colon`

Add a colon at end of atom

Type: function

Arity: 1 

Args: `(a)`


### Examples

```
> (colon (quote remember-this))
;;=>
remember-this:

```

-----------------------------------------------------
		

## `comma`

Add a comma at end of atom

Type: function

Arity: 1 

Args: `(a)`


### Examples

```
> (comma (quote hello))
;;=>
hello,

```

-----------------------------------------------------
		

## `comment`

Ignore the expressions in the block

Type: macro

Arity: 0+

Args: `(() . body)`


### Examples

```
> (comment twas brillig, and the slithy toves did gyre and gimble in the wabe)
;;=>
()

```

-----------------------------------------------------
		

## `complement`

Return the logical complement of the supplied function

Type: function

Arity: 1 

Args: `(f)`


### Examples

```
> ((complement even?) 1)
;;=>
t
> (map (complement odd?) (range 5))
;;=>
(t () t () t)

```

-----------------------------------------------------
		

## `concat`

Concatenenate any number of lists

Type: function

Arity: 0+

Args: `(() . lists)`


### Examples

```
> (concat (range 3) (quote (wow)) (reverse (range 3)))
;;=>
(0 1 2 wow 2 1 0)

```

-----------------------------------------------------
		

## `cond`

Fundamental branching construct

Type: special form

Arity: 0+

Args: `(() . pairs)`


### Examples

```
> (cond)
;;=> ()
> (cond (t 1) (t 2) (t 3))
;;=> 1
> (cond (() 1) (t 2))
;;=> 2

```

-----------------------------------------------------
		

## `cons`

Add an element to the front of a (possibly empty) list

Type: native function

Arity: 2 

Args: `(x xs)`


### Examples

```
> (cons 1 (quote (one two)))
;;=>
(1 one two)
> (cons 1 ())
;;=>
(1)
> (cons 1 2)
;;=>
(1 . 2)

```

-----------------------------------------------------
		

## `constantly`

Given a value, return a function which always returns that value

Type: function

Arity: 1 

Args: `(x)`


### Examples

```
> (map (constantly t) (range 10))
;;=>
(t t t t t t t t t t)

```

-----------------------------------------------------
		

## `dec`

Return the supplied integer argument, minus one

Type: function

Arity: 1 

Args: `(n)`


### Examples

```
> (dec 2)
;;=>
1
> (dec -1)
;;=>
-2

```

-----------------------------------------------------
		

## `def`

Set a value

Type: special form

Arity: 2 

Args: `(name value)`


-----------------------------------------------------
		

## `defmacro`

Create and name a macro

Type: special form

Arity: 2+

Args: `(name args . body)`


-----------------------------------------------------
		

## `defn`

Create and name a function

Type: special form

Arity: 2+

Args: `(name args . body)`


-----------------------------------------------------
		

## `doc`

Return the doclist for a function

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (doc (lambda (x) (doc (does stuff) (and other stuff)) (+ x 1)))
;;=>
((does stuff) (and other stuff))

```

-----------------------------------------------------
		

## `dotimes`

Execute body for each value in a list

Type: macro

Arity: 1+

Args: `(n . body)`


-----------------------------------------------------
		

## `downcase`

Return a new atom with all characters in lower case

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (downcase (quote Hello))
;;=>
hello
> (downcase (quote HELLO))
;;=>
hello

```

-----------------------------------------------------
		

## `drop`

Drop n items from a list, then return the rest

Type: function

Arity: 2 

Args: `(n l)`


### Examples

```
> (drop 3 (range 10))
;;=>
(3 4 5 6 7 8 9)

```

-----------------------------------------------------
		

## `error`

Raise an error

Type: special form

Arity: 1 

Args: `(msg-list)`


-----------------------------------------------------
		

## `errors`

Error checking (for tests)

Type: special form

Arity: 1+

Args: `(message-pattern-list . exprs)`


-----------------------------------------------------
		

## `even?`

Return true if the supplied integer argument is even

Type: function

Arity: 1 

Args: `(n)`


### Examples

```
> (map even? (range 5))
;;=>
(t () t () t)

```

-----------------------------------------------------
		

## `exclaimed`

Return l as a sentence... emphasized!

Type: function

Arity: 1 

Args: `(l)`


### Examples

```
> (exclaimed (quote (well, hello)))
;;=>
(Well, hello!)
> (exclaimed (quote (help)))
;;=>
(Help!)
> (exclaimed (quote (begone, fiend)))
;;=>
(Begone, fiend!)

```

-----------------------------------------------------
		

## `filter`

Keep only values for which function f is true

Type: function

Arity: 2 

Args: `(f l)`


### Examples

```
> (filter odd? (range 5))
;;=>
(1 3)

```

-----------------------------------------------------
		

## `flatten`

Return a (possibly nested) list, flattened

Type: function

Arity: 1 

Args: `(l)`


### Examples

```
> (flatten (quote (this is a (really (nested) list))))
;;=>
(this is a really nested list)

```

-----------------------------------------------------
		

## `foreach`

Execute body for each value in a list

Type: macro

Arity: 2+

Args: `(x xs . body)`


-----------------------------------------------------
		

## `forms`

Return available operators, as a list

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `fuse`

Fuse a list of numbers or atoms into a single atom

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (fuse (quote (A B C)))
;;=>
ABC
> (fuse (reverse (range 10)))
;;=>
9876543210

```

-----------------------------------------------------
		

## `help`

Print this message

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `identity`

Return the argument

Type: function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `if`

Simple conditional with two branches

Type: macro

Arity: 3 

Args: `(condition then else)`


-----------------------------------------------------
		

## `if-not`

Simple (inverted) conditional with two branches

Type: macro

Arity: 3 

Args: `(condition then else)`


-----------------------------------------------------
		

## `inc`

Return the supplied integer argument, plus one

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		

## `is`

Assert a condition is truthy, or show failing code

Type: macro

Arity: 1 

Args: `(expr)`


### Examples

```
> (is t)
;;=>
()
> (is (car (cons () (quote (this one should fail)))))
;;=>
ERROR: (assertion failed: (car (cons () (quote (this one should fail)))))

```

-----------------------------------------------------
		

## `lambda`

Create a function

Type: special form

Arity: 1+

Args: `(args . body) or (name args . body)`


-----------------------------------------------------
		

## `last`

Return the last item in a list

Type: function

Arity: 1 

Args: `(l)`


### Examples

```
> (last (range 10))
;;=>
9
> (last (split (quote ATOM!)))
;;=>
!

```

-----------------------------------------------------
		

## `len`

Return the length of a list

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (len (range 10))
;;=>
10

```

-----------------------------------------------------
		

## `let`

Create a local scope with bindings

Type: special form

Arity: 1+

Args: `(bindings . body)`


-----------------------------------------------------
		

## `list`

Return a list of the given arguments

Type: native function

Arity: 0+

Args: `(() . xs)`


### Examples

```
> (list 1 2 3)
;;=>
(1 2 3)
> (list)
;;=>
()

```

-----------------------------------------------------
		

## `list*`

Create a list by consing everything but the last arg onto the last

Type: function

Arity: 0+

Args: `(() . args)`


### Examples

```
> (list* 1 2 (quote (3)))
;;=>
(1 2 3)
> (list* 1 2 (quote (3 4)))
;;=>
(1 2 3 4)
> (list*)
;;=>
()

```

-----------------------------------------------------
		

## `list?`

Return t if the argument is a list, () otherwise

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (list? (range 10))
;;=>
t
> (list? 1)
;;=>
()

```

-----------------------------------------------------
		

## `loop`

Loop forever

Type: special form

Arity: 1+

Args: `(() . body)`


-----------------------------------------------------
		

## `macroexpand-1`

Expand a macro

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (macroexpand-1 (quote (+ x 1)))
;;=>
(+ x 1)
> (macroexpand-1 (quote (if () 1 2)))
;;=>
(cond (() 1) (t 2))

```

-----------------------------------------------------
		

## `map`

Apply the supplied function to every element in the supplied list

Type: function

Arity: 2 

Args: `(f l)`


### Examples

```
> (map odd? (range 5))
;;=>
(() t () t ())
> (map true? (quote (foo t () t 3)))
;;=>
(() t () t ())

```

-----------------------------------------------------
		

## `mapcat`

Map a function onto a list and concatenate results

Type: function

Arity: 2 

Args: `(f l)`


### Examples

```
> (map list (range 5))
;;=>
((0) (1) (2) (3) (4))
> (mapcat list (range 5))
;;=>
(0 1 2 3 4)
> (map range (range 5))
;;=>
(() (0) (0 1) (0 1 2) (0 1 2 3))
> (mapcat range (range 5))
;;=>
(0 0 1 0 1 2 0 1 2 3)

```

-----------------------------------------------------
		

## `neg?`

Return true iff the supplied integer argument is less than zero

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		

## `not`

Return t if the argument is nil, () otherwise

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (not ())
;;=>
t
> (not t)
;;=>
()
> (not (range 10))
;;=>
()

```

-----------------------------------------------------
		

## `nth`

Find the nth value of a list, starting from zero

Type: function

Arity: 2 

Args: `(n l)`


### Examples

```
> (nth 3 (quote (one two three four five)))
;;=>
four

```

-----------------------------------------------------
		

## `number?`

Return true if the argument is a number, else ()

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (number? 1)
;;=>
t
> (number? t)
;;=>
()
> (number? +)
;;=>
()

```

-----------------------------------------------------
		

## `odd?`

Return true if the supplied integer argument is odd

Type: function

Arity: 1 

Args: `(n)`


### Examples

```
> (map even? (range 5))
;;=>
(t () t () t)

```

-----------------------------------------------------
		

## `or`

Boolean or

Type: special form

Arity: 0+

Args: `(() . xs)`


### Examples

```
> (or)
;; => false
> (or t t)
;; => true
> (or t t ())
;; => t
```

-----------------------------------------------------
		

## `period`

Add a period at end of atom

Type: function

Arity: 1 

Args: `(a)`


### Examples

```
> (period (quote Woot))
;;=>
Woot.

```

-----------------------------------------------------
		

## `pos?`

Return true iff the supplied integer argument is greater than zero

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		

## `print`

Print the arguments

Type: native function

Arity: 0+

Args: `(() . xs)`


-----------------------------------------------------
		

## `printl`

Print a list argument, without parentheses

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `println`

Print the arguments and a newline

Type: native function

Arity: 0+

Args: `(() . xs)`


-----------------------------------------------------
		

## `progn`

Execute multiple statements, returning the last

Type: macro

Arity: 0+

Args: `(() . body)`


-----------------------------------------------------
		

## `punctuate`

Return x capitalized, with punctuation determined by the supplied function

Type: function

Arity: 2 

Args: `(f x)`


-----------------------------------------------------
		

## `punctuate-atom`

Add a punctuation mark at end of atom

Type: function

Arity: 2 

Args: `(a mark)`


### Examples

```
> (punctuate-atom (quote list) (quote *))
;;=>
list*
> (punctuate-atom (quote list) COLON)
;;=>
list:

```

-----------------------------------------------------
		

## `quote`

Quote an expression

Type: special form

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `randalpha`

Return a list of random (English/Latin/unaccented) alphabetic characters

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		

## `randchoice`

Return an element at random from the supplied list

Type: function

Arity: 1 

Args: `(l)`


-----------------------------------------------------
		

## `randigits`

Return a random integer between 0 and the argument minus 1

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		

## `randint`

Return a random integer between 0 and the argument minus 1

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `range`

List of integers from 0 to n

Type: function

Arity: 1 

Args: `(n)`


### Examples

```
> (range 10)
;;=>
(0 1 2 3 4 5 6 7 8 9)
> (len (range 100))
;;=>
100

```

-----------------------------------------------------
		

## `readlist`

Read a list from stdin

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `reduce`

Successively apply a function against a list of arguments

Type: function

Arity: 2+

Args: `(f x . args)`


-----------------------------------------------------
		

## `rem`

Return remainder when second arg divides first

Type: native function

Arity: 2 

Args: `(x y)`


-----------------------------------------------------
		

## `remove`

Keep only values for which function f is false / the empty list

Type: function

Arity: 2 

Args: `(f l)`


### Examples

```
> (remove odd? (range 5))
;;=>
(0 2 4)

```

-----------------------------------------------------
		

## `repeat`

Return a list of length n whose elements are all x

Type: function

Arity: 2 

Args: `(n x)`


### Examples

```
> (repeat 5 (quote repetitive))
;;=>
(repetitive repetitive repetitive repetitive repetitive)

```

-----------------------------------------------------
		

## `repeatedly`

Return a list of length n whose elements are made from calling f repeatedly

Type: function

Arity: 2 

Args: `(n f)`


### Examples

```
> (repeatedly 3 (lambda () (range 5)))
;;=>
((0 1 2 3 4) (0 1 2 3 4) (0 1 2 3 4))

```

-----------------------------------------------------
		

## `reverse`

Reverse a list

Type: function

Arity: 1 

Args: `(l)`


### Examples

```
> (= (quote (c b a)) (reverse (quote (a b c))))
;;=>
t

```

-----------------------------------------------------
		

## `screen-clear`

Clear the screen

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `screen-end`

Stop screen for text UIs, return to console mode

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `screen-get-key`

Return a keystroke as an atom

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `screen-size`

Return the screen size (width, height)

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `screen-start`

Start screen for text UIs

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `screen-write`

Write a string to the screen

Type: native function

Arity: 3 

Args: `(x y list)`


-----------------------------------------------------
		

## `shuffle`

Return a (quickly!) shuffled list

Type: native function

Arity: 1 

Args: `(xs)`


-----------------------------------------------------
		

## `sleep`

Sleep for the given number of milliseconds

Type: native function

Arity: 1 

Args: `(ms)`


-----------------------------------------------------
		

## `some`

Return f applied to first element for which that result is truthy, else ()

Type: function

Arity: 2 

Args: `(f l)`


### Examples

```
> (some even? (quote (1 3 5 7 9 11 13)))
;;=>
()
> (some even? (quote (1 3 5 7 9 1000 11 13)))
;;=>
t

```

-----------------------------------------------------
		

## `split`

Split an atom or number into a list of single-digit numbers or single-character atoms

Type: native function

Arity: 1 

Args: `(x)`


### Examples

```
> (split 123)
;;=>
(1 2 3)
> (split (quote abc))
;;=>
(a b c)

```

-----------------------------------------------------
		

## `syntax-quote`

Syntax-quote an expression

Type: special form

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `take`

Take up to n items from the supplied list

Type: function

Arity: 2 

Args: `(n l)`


### Examples

```
> (take 3 (range 10))
;;=>
(0 1 2)

```

-----------------------------------------------------
		

## `test`

Establish a testing block (return last expression)

Type: native function

Arity: 0+

Args: `(() . exprs)`


-----------------------------------------------------
		

## `tosentence`

Return l as a sentence... capitalized, with a period at the end

Type: function

Arity: 1 

Args: `(l)`


-----------------------------------------------------
		

## `true?`

Return t if the argument is t

Type: function

Arity: 1 

Args: `(x)`


### Examples

```
> (true? 3)
;;=>
()
> (true? t)
;;=>
t

```

-----------------------------------------------------
		

## `upcase`

Return the uppercase version of the given atom

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `version`

Return the version of the interpreter

Type: native function

Arity: 0 

Args: `()`


-----------------------------------------------------
		

## `when`

Simple conditional with single branch

Type: macro

Arity: 2 

Args: `(condition then)`


-----------------------------------------------------
		

## `when-not`

Complement of the when macro

Type: macro

Arity: 2 

Args: `(condition then)`


-----------------------------------------------------
		

## `with-screen`

Prepare for and clean up after screen operations

Type: macro

Arity: 0+

Args: `(() . body)`


-----------------------------------------------------
		

## `zero?`

Return true iff the supplied argument is zero

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		
