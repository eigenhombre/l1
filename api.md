107 forms available: <a href='#*'>*</a> <a href='#**'>**</a> <a href='#+'>+</a> <a href='#-'>-</a> <a href='#/'>/</a> <a href='#<'><</a> <a href='#<='><=</a> <a href='#='>=</a> <a href='#>'>></a> <a href='#>='>>=</a> <a href='#and'>and</a> <a href='#apply'>apply</a> <a href='#atom?'>atom?</a> <a href='#bang'>bang</a> <a href='#body'>body</a> <a href='#butlast'>butlast</a> <a href='#caar'>caar</a> <a href='#capitalize'>capitalize</a> <a href='#car'>car</a> <a href='#cdr'>cdr</a> <a href='#comma'>comma</a> <a href='#comment'>comment</a> <a href='#complement'>complement</a> <a href='#concat'>concat</a> <a href='#cond'>cond</a> <a href='#cons'>cons</a> <a href='#constantly'>constantly</a> <a href='#dec'>dec</a> <a href='#def'>def</a> <a href='#defmacro'>defmacro</a> <a href='#defn'>defn</a> <a href='#doc'>doc</a> <a href='#dotimes'>dotimes</a> <a href='#downcase'>downcase</a> <a href='#drop'>drop</a> <a href='#error'>error</a> <a href='#errors'>errors</a> <a href='#even?'>even?</a> <a href='#exclaimed'>exclaimed</a> <a href='#filter'>filter</a> <a href='#flatten'>flatten</a> <a href='#foreach'>foreach</a> <a href='#forms'>forms</a> <a href='#fuse'>fuse</a> <a href='#help'>help</a> <a href='#identity'>identity</a> <a href='#if'>if</a> <a href='#if-not'>if-not</a> <a href='#inc'>inc</a> <a href='#is'>is</a> <a href='#lambda'>lambda</a> <a href='#last'>last</a> <a href='#len'>len</a> <a href='#let'>let</a> <a href='#list'>list</a> <a href='#list*'>list*</a> <a href='#list?'>list?</a> <a href='#loop'>loop</a> <a href='#macroexpand-1'>macroexpand-1</a> <a href='#map'>map</a> <a href='#mapcat'>mapcat</a> <a href='#neg?'>neg?</a> <a href='#not'>not</a> <a href='#nth'>nth</a> <a href='#number?'>number?</a> <a href='#odd?'>odd?</a> <a href='#or'>or</a> <a href='#period'>period</a> <a href='#pos?'>pos?</a> <a href='#print'>print</a> <a href='#printl'>printl</a> <a href='#println'>println</a> <a href='#progn'>progn</a> <a href='#punctuate'>punctuate</a> <a href='#quote'>quote</a> <a href='#randalpha'>randalpha</a> <a href='#randchoice'>randchoice</a> <a href='#randigits'>randigits</a> <a href='#randint'>randint</a> <a href='#range'>range</a> <a href='#readlist'>readlist</a> <a href='#reduce'>reduce</a> <a href='#rem'>rem</a> <a href='#remove'>remove</a> <a href='#repeat'>repeat</a> <a href='#repeatedly'>repeatedly</a> <a href='#reverse'>reverse</a> <a href='#screen-clear'>screen-clear</a> <a href='#screen-end'>screen-end</a> <a href='#screen-get-key'>screen-get-key</a> <a href='#screen-size'>screen-size</a> <a href='#screen-start'>screen-start</a> <a href='#screen-write'>screen-write</a> <a href='#shuffle'>shuffle</a> <a href='#sleep'>sleep</a> <a href='#some'>some</a> <a href='#split'>split</a> <a href='#syntax-quote'>syntax-quote</a> <a href='#take'>take</a> <a href='#test'>test</a> <a href='#tosentence'>tosentence</a> <a href='#upcase'>upcase</a> <a href='#version'>version</a> <a href='#when'>when</a> <a href='#when-not'>when-not</a> <a href='#with-screen'>with-screen</a> <a href='#zero?'>zero?</a>

## `*`

Multiply 0 or more numbers

Type: native function

Arity: 0+

                 *  N    0+  Multiply 0 or more numbers

-----------------------------------------------------
		

## `**`

Exponentiation operator

Type: function

Arity: 2 

                **  F    2   Exponentiation operator

-----------------------------------------------------
		

## `+`

Add 0 or more numbers

Type: native function

Arity: 0+

                 +  N    0+  Add 0 or more numbers

-----------------------------------------------------
		

## `-`

Subtract 0 or more numbers from the first argument

Type: native function

Arity: 1+

                 -  N    1+  Subtract 0 or more numbers from the first argument

-----------------------------------------------------
		

## `/`

Divide the first argument by the rest

Type: native function

Arity: 2+

                 /  N    2+  Divide the first argument by the rest

-----------------------------------------------------
		

## `<`

Return t if the arguments are in strictly increasing order, () otherwise

Type: native function

Arity: 1+

                 <  N    1+  Return t if the arguments are in strictly increasing order, () otherwise

-----------------------------------------------------
		

## `<=`

Return t if the arguments are in increasing (or qual) order, () otherwise

Type: native function

Arity: 1+

                <=  N    1+  Return t if the arguments are in increasing (or qual) order, () otherwise

-----------------------------------------------------
		

## `=`

Return t if the arguments are equal, () otherwise

Type: native function

Arity: 1+

                 =  N    1+  Return t if the arguments are equal, () otherwise

-----------------------------------------------------
		

## `>`

Return t if the arguments are in strictly decreasing order, () otherwise

Type: native function

Arity: 1+

                 >  N    1+  Return t if the arguments are in strictly decreasing order, () otherwise

-----------------------------------------------------
		

## `>=`

Return t if the arguments are in decreasing (or equal) order, () otherwise

Type: native function

Arity: 1+

                >=  N    1+  Return t if the arguments are in decreasing (or equal) order, () otherwise

-----------------------------------------------------
		

## `and`

Boolean and

Type: special form

Arity: 0+

               and  S    0+  Boolean and

-----------------------------------------------------
		

## `apply`

Apply a function to a list of arguments

Type: native function

Arity: 2 

             apply  N    2   Apply a function to a list of arguments

-----------------------------------------------------
		

## `atom?`

Return t if the argument is an atom, () otherwise

Type: native function

Arity: 1 

             atom?  N    1   Return t if the argument is an atom, () otherwise

-----------------------------------------------------
		

## `bang`

Return a new atom with exclamation point added

Type: native function

Arity: 1 

              bang  N    1   Return a new atom with exclamation point added

-----------------------------------------------------
		

## `body`

Return the body of a lambda function

Type: native function

Arity: 1 

              body  N    1   Return the body of a lambda function

-----------------------------------------------------
		

## `butlast`

Return everything but the last element

Type: function

Arity: 1 

           butlast  F    1   Return everything but the last element

-----------------------------------------------------
		

## `caar`

First element of the first element of a list of lists

Type: function

Arity: 1 

              caar  F    1   First element of the first element of a list of lists

-----------------------------------------------------
		

## `capitalize`

Return the atom argument, capitalized

Type: function

Arity: 1 

        capitalize  F    1   Return the atom argument, capitalized

-----------------------------------------------------
		

## `car`

Return the first element of a list

Type: native function

Arity: 1 

               car  N    1   Return the first element of a list

-----------------------------------------------------
		

## `cdr`

Return a list with the first element removed

Type: native function

Arity: 1 

               cdr  N    1   Return a list with the first element removed

-----------------------------------------------------
		

## `comma`

Return a new atom with a comma at the end

Type: native function

Arity: 1 

             comma  N    1   Return a new atom with a comma at the end

-----------------------------------------------------
		

## `comment`

Ignore the expressions in the block

Type: macro

Arity: 0+

           comment  M    0+  Ignore the expressions in the block

-----------------------------------------------------
		

## `complement`

Return the logical complement of the supplied function

Type: function

Arity: 1 

        complement  F    1   Return the logical complement of the supplied function

-----------------------------------------------------
		

## `concat`

Concatenenate any number of lists

Type: function

Arity: 0+

            concat  F    0+  Concatenenate any number of lists

-----------------------------------------------------
		

## `cond`

Conditional branching

Type: special form

Arity: 0+

              cond  S    0+  Conditional branching

-----------------------------------------------------
		

## `cons`

Add an element to the front of a (possibly empty) list

Type: native function

Arity: 2 

              cons  N    2   Add an element to the front of a (possibly empty) list

-----------------------------------------------------
		

## `constantly`

Given a value, return a function which always returns that value

Type: function

Arity: 1 

        constantly  F    1   Given a value, return a function which always returns that value

-----------------------------------------------------
		

## `dec`

Return the supplied integer argument, minus one

Type: function

Arity: 1 

               dec  F    1   Return the supplied integer argument, minus one

-----------------------------------------------------
		

## `def`

Set a value

Type: special form

Arity: 2 

               def  S    2   Set a value

-----------------------------------------------------
		

## `defmacro`

Create and name a macro

Type: special form

Arity: 2+

          defmacro  S    2+  Create and name a macro

-----------------------------------------------------
		

## `defn`

Create and name a function

Type: special form

Arity: 2+

              defn  S    2+  Create and name a function

-----------------------------------------------------
		

## `doc`

Return the doclist for a function

Type: native function

Arity: 1 

               doc  N    1   Return the doclist for a function

-----------------------------------------------------
		

## `dotimes`

Execute body for each value in a list

Type: macro

Arity: 1+

           dotimes  M    1+  Execute body for each value in a list

-----------------------------------------------------
		

## `downcase`

Return a new atom with all characters in lower case

Type: native function

Arity: 1 

          downcase  N    1   Return a new atom with all characters in lower case

-----------------------------------------------------
		

## `drop`

Drop n items from a list, then return the rest

Type: function

Arity: 2 

              drop  F    2   Drop n items from a list, then return the rest

-----------------------------------------------------
		

## `error`

Return an error and (NOT IMPLEMENTED) short-circuit further processing

Type: function

Arity: 1 

             error  F    1   Return an error and (NOT IMPLEMENTED) short-circuit further processing

-----------------------------------------------------
		

## `errors`

Error checking (for tests)

Type: special form

Arity: 1+

            errors  S    1+  Error checking (for tests)

-----------------------------------------------------
		

## `even?`

Return true if the supplied integer argument is even

Type: function

Arity: 1 

             even?  F    1   Return true if the supplied integer argument is even

-----------------------------------------------------
		

## `exclaimed`

Return l as a sentence... emphasized!

Type: function

Arity: 1 

         exclaimed  F    1   Return l as a sentence... emphasized!

-----------------------------------------------------
		

## `filter`

Keep only values for which function f is true

Type: function

Arity: 2 

            filter  F    2   Keep only values for which function f is true

-----------------------------------------------------
		

## `flatten`

Return a (possibly nested) list, flattened

Type: function

Arity: 1 

           flatten  F    1   Return a (possibly nested) list, flattened

-----------------------------------------------------
		

## `foreach`

Execute body for each value in a list

Type: macro

Arity: 2+

           foreach  M    2+  Execute body for each value in a list

-----------------------------------------------------
		

## `forms`

Return available operators, as a list

Type: native function

Arity: 0 

             forms  N    0   Return available operators, as a list

-----------------------------------------------------
		

## `fuse`

Fuse a list of numbers or atoms into a single atom

Type: native function

Arity: 1 

              fuse  N    1   Fuse a list of numbers or atoms into a single atom

-----------------------------------------------------
		

## `help`

Print this message

Type: native function

Arity: 0 

              help  N    0   Print this message

-----------------------------------------------------
		

## `identity`

Return the argument

Type: function

Arity: 1 

          identity  F    1   Return the argument

-----------------------------------------------------
		

## `if`

Simple conditional with two branches

Type: macro

Arity: 3 

                if  M    3   Simple conditional with two branches

-----------------------------------------------------
		

## `if-not`

Simple (inverted) conditional with two branches

Type: macro

Arity: 3 

            if-not  M    3   Simple (inverted) conditional with two branches

-----------------------------------------------------
		

## `inc`

Return the supplied integer argument, plus one

Type: function

Arity: 1 

               inc  F    1   Return the supplied integer argument, plus one

-----------------------------------------------------
		

## `is`

Assert that the argument is truthy (not ())

Type: native function

Arity: 1 

                is  N    1   Assert that the argument is truthy (not ())

-----------------------------------------------------
		

## `lambda`

Create a function

Type: special form

Arity: 1+

            lambda  S    1+  Create a function

-----------------------------------------------------
		

## `last`

Return the last item in a list

Type: function

Arity: 1 

              last  F    1   Return the last item in a list

-----------------------------------------------------
		

## `len`

Return the length of a list

Type: native function

Arity: 1 

               len  N    1   Return the length of a list

-----------------------------------------------------
		

## `let`

Create a local scope

Type: special form

Arity: 1+

               let  S    1+  Create a local scope

-----------------------------------------------------
		

## `list`

Return a list of the given arguments

Type: native function

Arity: 0+

              list  N    0+  Return a list of the given arguments

-----------------------------------------------------
		

## `list*`

Create a list by consing everything but the last arg onto the last

Type: function

Arity: 0+

             list*  F    0+  Create a list by consing everything but the last arg onto the last

-----------------------------------------------------
		

## `list?`

Return t if the argument is a list, () otherwise

Type: native function

Arity: 1 

             list?  N    1   Return t if the argument is a list, () otherwise

-----------------------------------------------------
		

## `loop`

Loop forever

Type: special form

Arity: 1+

              loop  S    1+  Loop forever

-----------------------------------------------------
		

## `macroexpand-1`

Expand a macro

Type: native function

Arity: 1 

     macroexpand-1  N    1   Expand a macro

-----------------------------------------------------
		

## `map`

Apply the supplied function to every element in the supplied list

Type: function

Arity: 2 

               map  F    2   Apply the supplied function to every element in the supplied list

-----------------------------------------------------
		

## `mapcat`

Map a function onto a list and concatenate results

Type: function

Arity: 2 

            mapcat  F    2   Map a function onto a list and concatenate results

-----------------------------------------------------
		

## `neg?`

Return true iff the supplied integer argument is less than zero

Type: function

Arity: 1 

              neg?  F    1   Return true iff the supplied integer argument is less than zero

-----------------------------------------------------
		

## `not`

Return t if the argument is nil, () otherwise

Type: native function

Arity: 1 

               not  N    1   Return t if the argument is nil, () otherwise

-----------------------------------------------------
		

## `nth`

Find the nth value of a list, starting from zero

Type: function

Arity: 2 

               nth  F    2   Find the nth value of a list, starting from zero

-----------------------------------------------------
		

## `number?`

Return true if the argument is a number, else ()

Type: native function

Arity: 1 

           number?  N    1   Return true if the argument is a number, else ()

-----------------------------------------------------
		

## `odd?`

Return true if the supplied integer argument is odd

Type: function

Arity: 1 

              odd?  F    1   Return true if the supplied integer argument is odd

-----------------------------------------------------
		

## `or`

Boolean or

Type: special form

Arity: 0+

                or  S    0+  Boolean or

-----------------------------------------------------
		

## `period`

Return a new atom with a period added to the end

Type: native function

Arity: 1 

            period  N    1   Return a new atom with a period added to the end

-----------------------------------------------------
		

## `pos?`

Return true iff the supplied integer argument is greater than zero

Type: function

Arity: 1 

              pos?  F    1   Return true iff the supplied integer argument is greater than zero

-----------------------------------------------------
		

## `print`

Print the arguments

Type: native function

Arity: 0+

             print  N    0+  Print the arguments

-----------------------------------------------------
		

## `printl`

Print a list argument, without parentheses

Type: native function

Arity: 1 

            printl  N    1   Print a list argument, without parentheses

-----------------------------------------------------
		

## `println`

Print the arguments and a newline

Type: native function

Arity: 0+

           println  N    0+  Print the arguments and a newline

-----------------------------------------------------
		

## `progn`

Execute multiple statements, returning the last

Type: macro

Arity: 0+

             progn  M    0+  Execute multiple statements, returning the last

-----------------------------------------------------
		

## `punctuate`

Return l capitalized, with punctuation determined by the supplied function

Type: function

Arity: 2 

         punctuate  F    2   Return l capitalized, with punctuation determined by the supplied function

-----------------------------------------------------
		

## `quote`

Quote an expression

Type: special form

Arity: 1 

             quote  S    1   Quote an expression

-----------------------------------------------------
		

## `randalpha`

Return a list of random (English/Latin/unaccented) alphabetic characters

Type: function

Arity: 1 

         randalpha  F    1   Return a list of random (English/Latin/unaccented) alphabetic characters

-----------------------------------------------------
		

## `randchoice`

Return an element at random from the supplied list

Type: function

Arity: 1 

        randchoice  F    1   Return an element at random from the supplied list

-----------------------------------------------------
		

## `randigits`

Return a random integer between 0 and the argument minus 1

Type: function

Arity: 1 

         randigits  F    1   Return a random integer between 0 and the argument minus 1

-----------------------------------------------------
		

## `randint`

Return a random integer between 0 and the argument minus 1

Type: native function

Arity: 1 

           randint  N    1   Return a random integer between 0 and the argument minus 1

-----------------------------------------------------
		

## `range`

List of integers from 0 to n

Type: function

Arity: 1 

             range  F    1   List of integers from 0 to n

-----------------------------------------------------
		

## `readlist`

Read a list from stdin

Type: native function

Arity: 0 

          readlist  N    0   Read a list from stdin

-----------------------------------------------------
		

## `reduce`

Successively apply a function against a list of arguments

Type: function

Arity: 2+

            reduce  F    2+  Successively apply a function against a list of arguments

-----------------------------------------------------
		

## `rem`

Return remainder when second arg divides first

Type: native function

Arity: 2 

               rem  N    2   Return remainder when second arg divides first

-----------------------------------------------------
		

## `remove`

Keep only values for which function f is false / the empty list

Type: function

Arity: 2 

            remove  F    2   Keep only values for which function f is false / the empty list

-----------------------------------------------------
		

## `repeat`

Return a list of length n whose elements are all x

Type: function

Arity: 2 

            repeat  F    2   Return a list of length n whose elements are all x

-----------------------------------------------------
		

## `repeatedly`

Return a list of length n whose elements are made from calling f repeatedly

Type: function

Arity: 2 

        repeatedly  F    2   Return a list of length n whose elements are made from calling f repeatedly

-----------------------------------------------------
		

## `reverse`

Reverse a list

Type: function

Arity: 1 

           reverse  F    1   Reverse a list

-----------------------------------------------------
		

## `screen-clear`

Clear the screen

Type: native function

Arity: 0 

      screen-clear  N    0   Clear the screen

-----------------------------------------------------
		

## `screen-end`

Stop screen for text UIs, return to console mode

Type: native function

Arity: 0 

        screen-end  N    0   Stop screen for text UIs, return to console mode

-----------------------------------------------------
		

## `screen-get-key`

Return a keystroke as an atom

Type: native function

Arity: 0 

    screen-get-key  N    0   Return a keystroke as an atom

-----------------------------------------------------
		

## `screen-size`

Return the screen size (width, height)

Type: native function

Arity: 0 

       screen-size  N    0   Return the screen size (width, height)

-----------------------------------------------------
		

## `screen-start`

Start screen for text UIs

Type: native function

Arity: 0 

      screen-start  N    0   Start screen for text UIs

-----------------------------------------------------
		

## `screen-write`

Write a string to the screen

Type: native function

Arity: 3 

      screen-write  N    3   Write a string to the screen

-----------------------------------------------------
		

## `shuffle`

Return a shuffled list

Type: native function

Arity: 1 

           shuffle  N    1   Return a shuffled list

-----------------------------------------------------
		

## `sleep`

Sleep for the given number of milliseconds

Type: native function

Arity: 1 

             sleep  N    1   Sleep for the given number of milliseconds

-----------------------------------------------------
		

## `some`

Return f applied to first element for which that result is truthy, else ()

Type: function

Arity: 2 

              some  F    2   Return f applied to first element for which that result is truthy, else ()

-----------------------------------------------------
		

## `split`

Split an atom or number into a list of single-digit numbers or single-character atoms

Type: native function

Arity: 1 

             split  N    1   Split an atom or number into a list of single-digit numbers or single-character atoms

-----------------------------------------------------
		

## `syntax-quote`

Syntax-quote an expression

Type: special form

Arity: 1 

      syntax-quote  S    1   Syntax-quote an expression

-----------------------------------------------------
		

## `take`

Take up to n items from the supplied list

Type: function

Arity: 2 

              take  F    2   Take up to n items from the supplied list

-----------------------------------------------------
		

## `test`

Establish a testing block (return last expression)

Type: native function

Arity: 0+

              test  N    0+  Establish a testing block (return last expression)

-----------------------------------------------------
		

## `tosentence`

Return l as a sentence... capitalized, with a period at the end

Type: function

Arity: 1 

        tosentence  F    1   Return l as a sentence... capitalized, with a period at the end

-----------------------------------------------------
		

## `upcase`

Return the uppercase version of the given atom

Type: native function

Arity: 1 

            upcase  N    1   Return the uppercase version of the given atom

-----------------------------------------------------
		

## `version`

Return the version of the interpreter

Type: native function

Arity: 0 

           version  N    0   Return the version of the interpreter

-----------------------------------------------------
		

## `when`

Simple conditional with single branch

Type: macro

Arity: 2 

              when  M    2   Simple conditional with single branch

-----------------------------------------------------
		

## `when-not`

Complement of the when macro

Type: macro

Arity: 2 

          when-not  M    2   Complement of the when macro

-----------------------------------------------------
		

## `with-screen`

Prepare for and clean up after screen operations

Type: macro

Arity: 0+

       with-screen  M    0+  Prepare for and clean up after screen operations

-----------------------------------------------------
		

## `zero?`

Return true iff the supplied argument is zero

Type: function

Arity: 1 

             zero?  F    1   Return true iff the supplied argument is zero

-----------------------------------------------------
		
