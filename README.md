# l1

<img src="/l1-impression.jpg" width="400">

![build](https://github.com/eigenhombre/l1/actions/workflows/build.yml/badge.svg)

`l1` is a small [Lisp 1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace) written in Go:

- Numbers are "big" (Go's `big.Int`); integer math only, so far.
- `t` as True, `()` as `nil`;
- Side effects: assignment, e.g. `(def a 1234)`, and a `print` function;
- Lexical scope for functions.
- Simple error checking

| l1 has                                    | doesn't have | will have                 | might get                            |
| ----------------------------------------- |:------------:| :------------------------:|:-------------------------------------:
| ints (large)                              | keywords     | macros                    | curses                               |
| comments (;; ....)                        | maps         | syntax quote              | graphics                             |
| atoms                                     | strings      | reader macros (`, ', ...) | subprocess / shells                  |
| lists                                     | namespaces   | REPL / editor integration | big floats                           |
| 4 special forms: cond, def, lambda, quote | exceptions   | let (as a macro)          | rational numbers                     |
| 16 built-in functions                     | loops        | defun/defn (as a macro)   | tail call optimization               |
| recursion                                 |              |                           | `error` equivalent                   |
| closures                                  |              |                           | byte code compilation/interpretation |

# Usage / Examples

You should have Go installed and configured.  At some later point, pre-built
artifacts for various architectures may be available here.

## Installing Using the `go` Tool

    go install github.com/eigenhombre/l1@latest

## Building from Source

Check out this repo and `cd` to it. Then,

    go test
    go build
    go install

## Usage

To execute a file:

    $ l1 <file.l1>

Example (using a file in this project):

    $ cat fact.l1
    ;; Return the factorial of `n`:
    (def fact
         (lambda (n)
           (cond ((= 0 n) 1)
                 (t (* n (fact (- n 1)))))))

    (print (fact 100))
    $ time l1 fact.l1
    933262154439441526816992388562667004907159682643816214685929638
    952175999932299156089414639761565182862536979208272237582511852
    10916864000000000000000000000000

    real	0m0.008s
    user	0m0.004s
    sys	0m0.004s

## Example REPL Session

These were copied directly from the unit test output; `eval_test.go` has more examples.

<!-- BEGIN EXAMPLES --->
# l1

<img src="/l1-impression.jpg" width="400">

![build](https://github.com/eigenhombre/l1/actions/workflows/build.yml/badge.svg)

`l1` is a small [Lisp 1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace) written in Go:

- Numbers are "big" (Go's `big.Int`); integer math only, so far.
- `t` as True, `()` as `nil`;
- Side effects: assignment, e.g. `(def a 1234)`, and a `print` function;
- Lexical scope for functions.
- Simple error checking

| l1 has                                    | doesn't have | will have                 | might get                            |
| ----------------------------------------- |:------------:| :------------------------:|:-------------------------------------:
| ints (large)                              | keywords     | macros                    | curses                               |
| comments (;; ....)                        | maps         | syntax quote              | graphics                             |
| atoms                                     | strings      | reader macros (`, ', ...) | subprocess / shells                  |
| lists                                     | namespaces   | REPL / editor integration | big floats                           |
| 4 special forms: cond, def, lambda, quote | exceptions   | let (as a macro)          | rational numbers                     |
| 16 built-in functions                     | loops        | defun/defn (as a macro)   | tail call optimization               |
| recursion                                 |              |                           | `error` equivalent                   |
| closures                                  |              |                           | byte code compilation/interpretation |

# Usage / Examples

You should have Go installed and configured.  At some later point, pre-built
artifacts for various architectures may be available here.

## Installing Using the `go` Tool

    go install github.com/eigenhombre/l1@latest

## Building from Source

Check out this repo and `cd` to it. Then,

    go test
    go build
    go install

## Usage

To execute a file:

    $ l1 <file.l1>

Example (using a file in this project):

    $ cat fact.l1
    ;; Return the factorial of `n`:
    (def fact
         (lambda (n)
           (cond ((= 0 n) 1)
                 (t (* n (fact (- n 1)))))))

    (print (fact 100))
    $ time l1 fact.l1
    933262154439441526816992388562667004907159682643816214685929638
    952175999932299156089414639761565182862536979208272237582511852
    10916864000000000000000000000000

    real	0m0.008s
    user	0m0.004s
    sys	0m0.004s

## Example REPL Session

These were copied directly from the unit test output; `eval_test.go` has more examples.

<!-- BEGIN EXAMPLES --->


    $ l1
    > t
    t
    > (not t)
    ()
    > (not ())
    t
    > ()  ;; Nil by any other name, would still smell as sweet...
    ()
    > (cons t ())
    (t)
    > (cons (quote hello) (quote (world)))
    (hello world)
    > (quote foo)
    foo
    > (quote (the (ten (laws (of (greenspun))))))
    (the (ten (laws (of (greenspun)))))
    > (cdr (quote (is not common lisp)))
    (not common lisp)
    > (car (quote (is not common lisp)))
    is
    > (len (quote (1 2 3)))
    3
    > (+ 1 1)
    2
    > (+ 1 2)
    3
    > (* 12349807213490872130987 12349807213490872130987)
    152517738210391179737088822267441718485594169
    > (zero? 0)
    t
    > (zero? (quote zero))
    ()
    > (pos? 1)
    t
    > (neg? -1)
    t
    > (< 1 2 3)
    t
    > (<= 1 2 3 3)
    t
    > (> 3 2 1 0)
    t
    > (>= 3 2 1 1)
    t
    > (= (quote foo) (quote foo))
    t
    > (= (quote foo) (quote bar))
    ()
    > (= (quote foo) (quote (foo bar)))
    ()
    > (atom (quote (foo bar)))
    ()
    > (atom (quote atom))
    t
    > (cond (() 1) (2 3))
    3
    > (car (quote (1 2 3)))
    1
    > (cdr (quote (1 2 3)))
    (2 3)
    > (cons 1 (quote (2 3 4)))
    (1 2 3 4)
    > (split (quote greenspun))
    (g r e e n s p u n)
    > (split (* 12345 67890))
    (8 3 8 1 0 2 0 5 0)
    > (len (split (* 99999 99999 99999)))
    15
    > (fuse (quote (a b)))
    ab
    > (+ 2 (fuse (quote (1 2 3))))
    125
    > (fuse (split 1295807125987))
    1295807125987
    > (len (randigits 10))
    10
    > (apply + (quote (1 2 3)))
    6
    > (apply * (split 123456789))
    362880
    > (apply / (split 1111))
    1
    > (apply = (split (quote ooo)))
    t
    > (apply = (split (quote foo)))
    ()
    > (apply (lambda (x y z) (= x y z)) (split 121))
    ()
    > ((cond (t +)))
    0
    > ((car (cons + ())) 1 2 3)
    6
    > (def a 6)
    6
    > (def b 7)
    7
    > (+ a b)
    13
    > ((lambda ()))
    ()
    > ((lambda (x) (+ 1 x)) 1)
    2
    > (def fact (lambda (n) (cond ((= 0 n) 1) (t (* n (fact (- n 1)))))))
    <lambda(n)>
    > (fact 50)
    30414093201713378043612608166064768844377641568960512000000000000
    > (len (split (fact 1000)))
    2568
    > (def fib (lambda (n) (cond ((= 0 n) 0) ((= 1 n) 1) (t (+ (fib (- n 1)) (fib (- n 2)))))))
    <lambda(n)>
    > (fib 0)
    0
    > (fib 1)
    1
    > (fib 7)
    13
    > (fib 10)
    55
    > (fib 20)
    6765
    > (def a 1)
    1
    > (def f (lambda () (def a 2) a))
    <lambda()>
    > (f)
    2
    > a
    1
    > (def incrementer (lambda (n) (lambda (x) (+ x n))))
    <lambda(n)>
    > (def inc (incrementer 1))
    <lambda(x)>
    > (inc 5)
    6
    > (def add2 (incrementer 2))
    <lambda(x)>
    > (add2 5)
    7
    > (help)
    Builtins and Special Forms:
          Name  Arity    Description
             *    0+     Multiply 0 or more numbers
             +    0+     Add 0 or more numbers
             -    1+     Subtract 0 or more numbers from the first argument
             /    2+     Divide the first argument by the rest
             <    1+     Return t if the arguments are in strictly increasing order, () otherwise
            <=    1+     Return t if the arguments are in increasing (or qual) order, () otherwise
             =    1+     Return t if the arguments are equal, () otherwise
             >    1+     Return t if the arguments are in strictly decreasing order, () otherwise
            >=    1+     Return t if the arguments are in decreasing (or equal) order, () otherwise
         apply    2      Apply a function to a list of arguments
          atom    1      Return true if the argument is an atom, false otherwise
           car    1      Return the first element of a list
           cdr    1      Return a list with the first element removed
          cond    0+     SPECIAL FORM: Conditional branching
          cons    2      Add an element to the front of a (possibly empty) list
           def    2      SPECIAL FORM: Set a value
          fuse    1      Fuse a list of numbers or atoms into a single atom
          help    0      Print this message
        lambda    1+     SPECIAL FORM: Create a function
           len    1      Return the length of a list
          neg?    1      Return true if the (numeric) argument is negative, else ()
           not    1      Return t if the argument is nil, () otherwise
          pos?    1      Return true if the (numeric) argument is positive, else ()
         print    0+     Print the arguments
         quote    1      SPECIAL FORM: Quote an expression
     randigits    1      Return a list of random digits of the given length
         split    0      Split an atom or number into a list of single-digit numbers or single-character atoms
         zero?    1      Return t if the argument is zero, () otherwise
    > ^D
    $
<!--- END EXAMPLES -->

# CI/CD

A `Makefile` exists for convenience (combining testing, linting and build), and a `Dockerfile` is  used by a GitHub action for this project to email an alert if code is pushed which fails the build.

# Goals

- Learn more about Lisp as a model for computation by building a Lisp with sufficient power to [implement itself](http://www.paulgraham.com/rootsoflisp.html);
- Improve my Go skills;
- Build a small, fast-loading Lisp that I can extend how I like;
- Possibly implement Curses-based terminal control for text games, command line utilities, ...;

# Non-goals

- Backwards compatibility with any existing, popular Lisp.
- Stability (for now) -- everything is subject to change.

# Resources / Further Reading

- [Structure and Interpretation of Computer Programs](https://mitpress.mit.edu/sites/default/files/sicp/index.html).  Classic MIT
  text, presents several Lisp evaluation models, written in Scheme.
- [Crafting Interpreters](https://craftinginterpreters.com/) book / website.  Stunning, thorough,
  approachable and beautiful book on building a language in Java and
  in C.
- Donovan & Kernighan, [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440). Great Go reference.
- Rob Pike, [Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE) (YouTube).  I took the code described in this talk and spun it out into [its own package](https://github.com/eigenhombre/lexutil/) for reuse in `l1`.
- A [more detailed blog post](http://johnj.com/posts/l1/) on `l1`.

# License

Copyright © 2022, John Jacobsen. MIT License.

# Disclaimer

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
