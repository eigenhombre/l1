# API Index
109 forms available:
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
(* 1 2 3)
;;=>
6
(*)
;;=>
1

```

-----------------------------------------------------
		

## `**`

Exponentiation operator

Type: function

Arity: 2 

Args: `(n m)`


-----------------------------------------------------
		

## `+`

Add 0 or more numbers

Type: native function

Arity: 0+

Args: `(() . xs)`


-----------------------------------------------------
		

## `-`

Subtract 0 or more numbers from the first argument

Type: native function

Arity: 1+

Args: `(x . xs)`


-----------------------------------------------------
		

## `/`

Divide the first argument by the rest

Type: native function

Arity: 2+

Args: `(x . xs)`


-----------------------------------------------------
		

## `<`

Return t if the arguments are in strictly increasing order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


-----------------------------------------------------
		

## `<=`

Return t if the arguments are in increasing (or qual) order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


-----------------------------------------------------
		

## `=`

Return t if the arguments are equal, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


-----------------------------------------------------
		

## `>`

Return t if the arguments are in strictly decreasing order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


-----------------------------------------------------
		

## `>=`

Return t if the arguments are in decreasing (or equal) order, () otherwise

Type: native function

Arity: 1+

Args: `(x . xs)`


-----------------------------------------------------
		

## `and`

Boolean and

Type: special form

Arity: 0+

Args: `(() . xs)`


### Examples
```
(and)
;; => true
(and t t)
;; => true
(and t t ())
;; => ()

```

-----------------------------------------------------
		

## `apply`

Apply a function to a list of arguments

Type: native function

Arity: 2 

Args: `(f args)`


-----------------------------------------------------
		

## `atom?`

Return t if the argument is an atom, () otherwise

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `bang`

Add an exclamation point at end of atom

Type: function

Arity: 1 

Args: `(a)`


-----------------------------------------------------
		

## `body`

Return the body of a lambda function

Type: native function

Arity: 1 

Args: `(f)`


-----------------------------------------------------
		

## `butlast`

Return everything but the last element

Type: function

Arity: 1 

Args: `(l)`


-----------------------------------------------------
		

## `caar`

First element of the first element of a list of lists

Type: function

Arity: 1 

Args: `(l)`


-----------------------------------------------------
		

## `capitalize`

Return the atom argument, capitalized

Type: function

Arity: 1 

Args: `(a)`


-----------------------------------------------------
		

## `car`

Return the first element of a list

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `cdr`

Return a list with the first element removed

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `colon`

Add a colon at end of atom

Type: function

Arity: 1 

Args: `(a)`


-----------------------------------------------------
		

## `comma`

Add a comma at end of atom

Type: function

Arity: 1 

Args: `(a)`


-----------------------------------------------------
		

## `comment`

Ignore the expressions in the block

Type: macro

Arity: 0+

Args: `(() . body)`


-----------------------------------------------------
		

## `complement`

Return the logical complement of the supplied function

Type: function

Arity: 1 

Args: `(f)`


-----------------------------------------------------
		

## `concat`

Concatenenate any number of lists

Type: function

Arity: 0+

Args: `(() . lists)`


-----------------------------------------------------
		

## `cond`

Fundamental branching construct

Type: special form

Arity: 0+

Args: `(() . pairs)`


### Examples
```
(cond)
;; => ()
(cond (t 1) (t 2) (t 3))
;; => 1
(cond (() 1) (t 2))
;; => 2

```

-----------------------------------------------------
		

## `cons`

Add an element to the front of a (possibly empty) list

Type: native function

Arity: 2 

Args: `(x xs)`


-----------------------------------------------------
		

## `constantly`

Given a value, return a function which always returns that value

Type: function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `dec`

Return the supplied integer argument, minus one

Type: function

Arity: 1 

Args: `(n)`


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


-----------------------------------------------------
		

## `drop`

Drop n items from a list, then return the rest

Type: function

Arity: 2 

Args: `(n l)`


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


-----------------------------------------------------
		

## `exclaimed`

Return l as a sentence... emphasized!

Type: function

Arity: 1 

Args: `(l)`


-----------------------------------------------------
		

## `filter`

Keep only values for which function f is true

Type: function

Arity: 2 

Args: `(f l)`


-----------------------------------------------------
		

## `flatten`

Return a (possibly nested) list, flattened

Type: function

Arity: 1 

Args: `(l)`


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


-----------------------------------------------------
		

## `len`

Return the length of a list

Type: native function

Arity: 1 

Args: `(x)`


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


-----------------------------------------------------
		

## `list*`

Create a list by consing everything but the last arg onto the last

Type: function

Arity: 0+

Args: `(() . args)`


-----------------------------------------------------
		

## `list?`

Return t if the argument is a list, () otherwise

Type: native function

Arity: 1 

Args: `(x)`


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


-----------------------------------------------------
		

## `map`

Apply the supplied function to every element in the supplied list

Type: function

Arity: 2 

Args: `(f l)`


-----------------------------------------------------
		

## `mapcat`

Map a function onto a list and concatenate results

Type: function

Arity: 2 

Args: `(f l)`


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


-----------------------------------------------------
		

## `nth`

Find the nth value of a list, starting from zero

Type: function

Arity: 2 

Args: `(n l)`


-----------------------------------------------------
		

## `number?`

Return true if the argument is a number, else ()

Type: native function

Arity: 1 

Args: `(x)`


-----------------------------------------------------
		

## `odd?`

Return true if the supplied integer argument is odd

Type: function

Arity: 1 

Args: `(n)`


-----------------------------------------------------
		

## `or`

Boolean or

Type: special form

Arity: 0+

Args: `(() . xs)`


### Examples
```
(or)
;; => false
(or t t)
;; => true
(or t t ())
;; => t
```

-----------------------------------------------------
		

## `period`

Add a period at end of atom

Type: function

Arity: 1 

Args: `(a)`


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


-----------------------------------------------------
		

## `repeat`

Return a list of length n whose elements are all x

Type: function

Arity: 2 

Args: `(n x)`


-----------------------------------------------------
		

## `repeatedly`

Return a list of length n whose elements are made from calling f repeatedly

Type: function

Arity: 2 

Args: `(n f)`


-----------------------------------------------------
		

## `reverse`

Reverse a list

Type: function

Arity: 1 

Args: `(l)`


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


-----------------------------------------------------
		

## `split`

Split an atom or number into a list of single-digit numbers or single-character atoms

Type: native function

Arity: 1 

Args: `(x)`


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
		
