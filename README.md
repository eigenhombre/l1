# l1

<img src="/l1.jpg" width="400">

![build](https://github.com/eigenhombre/l1/actions/workflows/build.yml/badge.svg)

`l1` is a small interpreted [Lisp
1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace)
written in Go.  Emphasizing simplicity of data types (atoms,
arbitrarily large integers, and lists) and start-up speed, it aims to
be a playground for learning Lisp, making simple games, and exploring
classic (pre-ML) AI work.

`l1` eschews strings, vectors, and records in favor of using atoms and
lists in the style of [some classic AI
programs](https://github.com/norvig/paip-lisp).  It features macros,
tail-call optimization, unicode support, and a few unique functions
for converting atoms and numbers to lists, and vice-versa.

# Objectives

1. Provide a Lisp one can easily build
   [roguelikes](https://github.com/eigenhombre/onomat/) and other text
   games with.
1. Provide a Lisp for studying "[Good Old Fashioned
   AI](https://en.wikipedia.org/wiki/Symbolic_artificial_intelligence)"
   programs such as those described in [Paradigms of Artificial
   Intelligence Programming](https://github.com/norvig/paip-lisp)
   (PAIP).
1. Answer the question, How far can we get without a 'string' data
   type?  (Much older Lisp pedagogy, including PAIP and
   [SICP](https://en.wikipedia.org/wiki/Structure_and_Interpretation_of_Computer_Programs),
   does not refer to strings much or at all.)
1. Build a minimal language core and implement much of the language in
   the language itself (as is done in Clojure and many Common Lisp
   implementations).

# Documentation

See the [language
reference](https://github.com/eigenhombre/l1/blob/master/l1.md) and
the [API
docs](https://github.com/eigenhombre/l1/blob/master/l1.md#api-index).

# Status

The language core is largely complete, though things could still
change... *caveat emptor.*  Issues and remaining work are [tracked on
GitHub](https://github.com/eigenhombre/l1/issues).

Although there are many tests, I expect some bugs remain.  The
interpreter starts fast but, like most tree-walking interpreters, is
slow for longer calculations.

# Setup

## Option 1: Install using Go

To use this method, have Go installed and configured (including [setting
GOPATH](https://medium.com/@devesu/golang-quickstart-with-homebrew-macos-f3b3dacbc5dc)).

To install `l1`,

    go install github.com/eigenhombre/l1@latest

Specific versions are tagged and available as well.  See [the tags
page](https://github.com/eigenhombre/l1/tags) for available versions
and then, for example,

    go install github.com/eigenhombre/l1@v0.0.42

## Option 2: Download Pre-built Binary

If you don't want to install Go, `l1` binaries are built for 22 OS /
architecture combinations for every release.  (WARNING: many/most of
these are not tested regularly!)

1. Head on over to the [Releases Page](https://github.com/eigenhombre/l1/releases).
1. Download the binary appropriate for your operating system.
1. Rename it to a file `l1`.
1. Make `l1` executable.  On Unix / Linux / MacOS sytems, `chmod +x l1`
1. Move it to somewhere on your path (this is platform-specific).

### Note for Mac users:

Mac binaries are currently unsigned, so you have to go through the following extra steps:

1. Perform the above actions to get an executable file `l1` on your path.
1. Right-click (or control-click) the executable in the Finder.
1. Open with Terminal or iTerm or... whatever Terminal program you usually use.
1. Select "Open Anyway" after acknowledging the security warning.

You should only have to do this once per release, in general.

# Usage

At this point you should be able to run `l1`:

    $ l1 -h

to print a help message, or, to start a REPL:

    $ l1

To execute a file:

    $ l1 <file.l1>

Example, using a file in this project:

    $ cat examples/fact.l1
    ;; Return the factorial of `n`:
    (defn fact (n)
      (if (zero? n)
        1
        (* n (fact (- n 1)))))

    (print (fact 100))
    $ time l1 examples/fact.l1
    933262154439441526816992388562667004907159682643816214685929638
    952175999932299156089414639761565182862536979208272237582511852
    10916864000000000000000000000000

    real	0m0.012s
    user	0m0.007s
    sys	0m0.005s

See [these
instructions](https://github.com/eigenhombre/l1/blob/master/l1.md#making-binary-executables)
for turning an `l1` program into an executable binary.

# Example Programs

Several example programs are available in the
[`examples/`](https://github.com/eigenhombre/l1/tree/master/examples)
directory.  Most of these are run automatically as tests in the build.
These can be obtained either by cloning this repository, or if you
installed `l1` using the Go installer, by looking under `$GOPATH`. For
example, if the `latest` release is `v0.0.42`,

    $ ls $GOPATH/pkg/mod/github.com/eigenhombre/l1\@v0.0.42/examples/

[An example roguelike game](https://github.com/eigenhombre/onomat/)
lives in [a separate
repository](https://github.com/eigenhombre/onomat/).

## Example REPL Session

<!-- The following examples are autogenerated, do not change by hand! -->
<!-- BEGIN EXAMPLES -->

    $ l1
    > (quote foo)
    foo
    > 'foo
    foo
    > (len (split '水果))
    2
    > (quote (the (ten (laws (of (greenspun))))))
    (the (ten (laws (of (greenspun)))))
    > (cadaaaaaaaaaar '(((((((((((hello world))))))))))))
    world
    > ((lambda (x . xs) (list x xs)) 1 2 3 4)
    (1 (2 3 4))
    > (help)
    l1 - a Lisp interpreter.
    
                   Type
                   ---
                    S - special form
                    M - macro
                    N - native (Go) function
                    F - Lisp function
    
              Name Type Arity  Description
              ---- ---  ----  -----------
                 *  N    0+  Multiply 0 or more numbers
                **  F    2   Exponentiation operator
                 +  N    0+  Add 0 or more numbers
                 -  N    1+  Subtract 0 or more numbers from the first argument
                 /  N    2+  Divide the first argument by the rest
                 <  N    1+  Return t if the arguments are in strictly increasing order, () otherwise
                <=  N    1+  Return t if the arguments are in increasing or equal order, () otherwise
                 =  N    1+  Return t if the arguments are equal, () otherwise
                 >  N    1+  Return t if the arguments are in strictly decreasing order, () otherwise
                >=  N    1+  Return t if the arguments are in decreasing or equal order, () otherwise
               abs  F    1   Return absolute value of x
               and  S    0+  Boolean and
             apply  N    2   Apply a function to a list of arguments
             atom?  N    1   Return t if the argument is an atom, () otherwise
              bang  F    1   Add an exclamation point at end of atom
              body  N    1   Return the body of a lambda function
           butlast  F    1   Return everything but the last element
        capitalize  F    1   Return the atom argument, capitalized
               car  N    1   Return the first element of a list
               cdr  N    1   Return a list with the first element removed
             colon  F    1   Add a colon at end of atom
             comma  F    1   Add a comma at end of atom
           comment  M    0+  Ignore the expressions in the block
              comp  F    0+  Function composition -- return a function which applies a series of functions in reverse order
        complement  F    1   Return the logical complement of the supplied function
            concat  F    0+  Concatenenate any number of lists
           concat2  F    2   Concatenate two lists
              cond  S    0+  Fundamental branching construct
              cons  N    2   Add an element to the front of a (possibly empty) list
        constantly  F    1   Given a value, return a function which always returns that value
               dec  F    1   Return the supplied integer argument, minus one
               def  S    2   Set a value
          defmacro  S    2+  Create and name a macro
              defn  S    2+  Create and name a function
               doc  N    1   Return the doclist for a function
           dotimes  M    1+  Execute body for each value in a list
          downcase  N    1   Return a new atom with all characters in lower case
              drop  F    2   Drop n items from a list, then return the rest
         enumerate  F    1   Returning list of (i, x) pairs where i is the index (from zero) and x is the original element from l
             error  S    1   Raise an error
            errors  S    1+  Error checking, for tests
              eval  N    1   Evaluate an expression
             even?  F    1   Return true if the supplied integer argument is even
             every  F    2   Return t if f applied to every element in l is truthy, else ()
           exclaim  F    1   Return l as a sentence... emphasized!
              exit  N    0   Exit the program
            filter  F    2   Keep only values for which function f is true
           flatten  F    1   Return a (possibly nested) list, flattened
           foreach  M    2+  Execute body for each value in a list
             forms  N    0   Return available operators, as a list
              fuse  N    1   Fuse a list of numbers or atoms into a single atom
            gensym  N    0+  Return a new symbol
              help  N    0   Print a help message
          identity  F    1   Return the argument
                if  M    3   Simple conditional with two branches
            if-not  M    3   Simple (inverted) conditional with two branches
               inc  F    1   Return the supplied integer argument, plus one
         interpose  F    2   Interpose x between all elements of l
                is  M    1   Assert a condition is truthy, or show failing code
             isqrt  N    1   Integer square root
              juxt  F    0+  Create a function which combines multiple operations into a single list of results
            lambda  S    1+  Create a function
              last  F    1   Return the last item in a list
               len  N    1   Return the length of a list
               let  S    1+  Create a local scope with bindings
              let*  M    1+  Let form with ability to refer to previously-bound pairs in the binding list
              list  N    0+  Return a list of the given arguments
             list*  F    0+  Create a list by consing everything but the last arg onto the last
             list?  N    1   Return t if the argument is a list, () otherwise
              load  N    1   Load and execute a file
              loop  S    1+  Loop forever
     macroexpand-1  N    1   Expand a macro
               map  F    2   Apply the supplied function to every element in the supplied list
            mapcat  F    2   Map a function onto a list and concatenate results
               max  F    0+  Find maximum of one or more numbers
               min  F    0+  Find minimum of one or more numbers
              neg?  F    1   Return true iff the supplied integer argument is less than zero
               not  N    1   Return t if the argument is nil, () otherwise
              not=  F    0+  Complement of = function
               nth  F    2   Find the nth value of a list, starting from zero
           number?  N    1   Return true if the argument is a number, else ()
              odd?  F    1   Return true if the supplied integer argument is odd
                or  S    0+  Boolean or
           partial  F    1+  Partial function application
            period  F    1   Add a period at end of atom
              pos?  F    1   Return true iff the supplied integer argument is greater than zero
             print  N    0+  Print the arguments
            printl  N    1   Print a list argument, without parentheses
           println  N    0+  Print the arguments and a newline
             progn  M    0+  Execute multiple statements, returning the last
         punctuate  F    2   Return x capitalized, with punctuation determined by the supplied function
    punctuate-atom  F    2   Add a punctuation mark at end of atom
             quote  S    1   Quote an expression
         randalpha  F    1   Return a list of random (English/Latin/unaccented) lower-case alphabetic characters
        randchoice  F    1   Return an element at random from the supplied list
         randigits  F    1   Return a random integer between 0 and the argument minus 1
           randint  N    1   Return a random integer between 0 and the argument minus 1
             range  F    1   List of integers from 0 to n
          readlist  N    0   Read a list from stdin
            reduce  F    2+  Successively apply a function against a list of arguments
               rem  N    2   Return remainder when second arg divides first
            remove  F    2   Keep only values for which function f is false / the empty list
            repeat  F    2   Return a list of length n whose elements are all x
        repeatedly  F    2   Return a list of length n whose elements are made from calling f repeatedly
           reverse  F    1   Reverse a list
      screen-clear  N    0   Clear the screen
        screen-end  N    0   Stop screen for text UIs, return to console mode
    screen-get-key  N    0   Return a keystroke as an atom
       screen-size  N    0   Return the screen size: width, height
      screen-start  N    0   Start screen for text UIs
      screen-write  N    3   Write a string to the screen
            second  F    1   Return the second element of a list, or () if not enough elements
              set!  S    2   Update a value in an existing binding
             shell  N    1   Run a shell subprocess, and return stdout, stderr, and exit code
           shuffle  N    1   Return a (quickly!) shuffled list
             sleep  N    1   Sleep for the given number of milliseconds
              some  F    2   Return f applied to first element for which that result is truthy, else ()
              sort  N    1   Sort a list
           sort-by  N    2   Sort a list by a function
            source  N    1   Show source for a function
             split  N    1   Split an atom or number into a list of single-digit numbers or single-character atoms
           swallow  S    0+  Swallow errors thrown in body, return t if any occur
      syntax-quote  S    1   Syntax-quote an expression
              take  F    2   Take up to n items from the supplied list
              test  S    0+  Run tests
        tosentence  F    1   Return l as a sentence... capitalized, with a period at the end
             true?  F    1   Return t if the argument is t
               try  S    0+  Try to evaluate body, catch errors and handle them
            upcase  N    1   Return the uppercase version of the given atom
           version  N    0   Return the version of the interpreter
              when  M    1+  Simple conditional with single branch
          when-not  M    1+  Complement of the when macro
             while  M    1+  Loop for as long as condition is true
       with-screen  M    0+  Prepare for and clean up after screen operations
             zero?  F    1   Return true iff the supplied argument is zero
    > ^D
    $
<!-- END EXAMPLES -->

Many of the [unit tests](https://github.com/eigenhombre/l1/blob/master/tests.l1) are written in `l1` itself.  Here are a few examples:

```
(test '(split and fuse)
  (is (= '(1) (split 1)))
  (is (= '(-1) (split -1)))
  (is (= '(-3 2 1) (split -321)))
  (is (= '(a) (split (quote a))))
  (is (= '(g r e e n s p u n) (split 'greenspun)))
  (is (= '(8 3 8 1 0 2 0 5 0) (split (* 12345 67890))))
  (is (= 15 (len (split (* 99999 99999 99999)))))
  (errors '(expects a single argument)
    (split))
  (errors '(expects a single argument)
    (split 1 1))
  (errors '(expects an atom or a number)
    (split '(a b c)))

  (is (= '() (fuse ())))
  (is (= 'a (fuse (quote (a)))))
  (is (= 'aa (fuse (quote (aa)))))
  (is (= 'ab (fuse (quote (a b)))))
  (is (= 1 (fuse (quote (1)))))
  (is (= 12 (fuse (quote (1 2)))))
  (is (= 125 (+ 2 (fuse (quote (1 2 3))))))
  (is (= 1295807125987 (fuse (split 1295807125987))))
  (errors '(expects a single argument)
    (fuse)))

(test '(factorial)
  (defn fact (n)
    (if (zero? n)
      1
      (* n (fact (- n 1))))))
  (is (= 30414093201713378043612608166064768844377641568960512000000000000
         (fact 50)))
  (is (= 2568 (len (split (fact 1000))))))
```

Several core library functions are also implemented in `l1`.  The file
[`l1.l1`](https://github.com/eigenhombre/l1/blob/master/lisp/l1.l1)
contains these, and is evaluated when the interpreter starts.  This
currently runs quite quickly (about 12 milliseconds on my Mac M1 Air).

# Local Development

Check out this repo and `cd` to it. Then,

    go test
    go build
    go install

Or, for a complete build with all tests, etc.,

    make

Testing and builds rely on GitHub Actions, Docker, and Make.  Please
look at the `Dockerfile` and `Makefile` for more information.

New releases are made using `make release`.  You must commit all
outstanding changes first.

# Resources / Further Reading

- [Structure and Interpretation of Computer
  Programs](https://en.wikipedia.org/wiki/Structure_and_Interpretation_of_Computer_Programs).
  Classic MIT text, presents several Lisp evaluation models, written
  in Scheme.
- [Crafting Interpreters](https://craftinginterpreters.com/) book / website.  Stunning, thorough,
  approachable and beautiful book on building a language in Java and
  in C.
- Donovan & Kernighan, [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440). Great Go reference.
- Rob Pike, [Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE) (YouTube).  I took the code described in this talk and spun it out into [its own package](https://github.com/eigenhombre/lexutil/) for reuse in `l1`.
- A [more detailed blog post](http://johnj.com/posts/l1/) on `l1`.
- A [blog post on adding Tail Call Optimization](http://johnj.com/posts/tco/) to `l1`.

# License

Copyright © 2022-2024, John Jacobsen. MIT License.

Two of the programs in `examples/` were adapted from P. Norvig,
[Paradigms of Artificial Intelligence Programming: Case Studies in
Common Lisp](https://github.com/norvig/paip-lisp). MIT License.

# Disclaimer

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
