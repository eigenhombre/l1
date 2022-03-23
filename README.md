# l1

<img src="/l1.jpg" width="400">

![build](https://github.com/eigenhombre/l1/actions/workflows/build.yml/badge.svg)

`l1` is a small Lisp written in Go:

- [Lisp 1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace);
- Numbers are "big" (Go's `big.Int`); integer math only, so far.
- Eval:
  - `t` as True, `()` as `nil`;
  - Arithmetical operators `+ - * /`;
  - `quote`, `car`, `cdr`, `cons`, `cond` and a few other built-ins;
  - Side effects: assignment, e.g. `(def a 1234)`, and a `print` function;
  - Atoms bind to values in the local or global environment;
  - Lexical scope for functions.
- Simple error checking

# Usage / Examples

## Installing Using the `go` Tool

    go get github.com/eigenhombre/l1@latest

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
           (cond ((eq 0 n) 1) 
                 (t (* n (fact (- n 1)))))))

    (print (fact 100))
    $ time l1 fact.l1 
    933262154439441526816992388562667004907159682643816214685929638
    952175999932299156089414639761565182862536979208272237582511852
    10916864000000000000000000000000

    real	0m0.008s
    user	0m0.004s
    sys	0m0.004s

Compare with [Babashka](https://github.com/babashka/babashka):

    $ cat fact.clj
    ;; Return the factorial of `n`:
    (def fact
      (fn [n]
        (cond (= 0 n) 1
              :else (*' n (fact (- n 1))))))

    (println (fact 100))    
    $ time bb fact.clj
    933262154439441526816992388562667004907159682643816214685929638
    952175999932299156089414639761565182862536979208272237582511852
    10916864000000000000000000000000N

    real	0m0.200s
    user	0m0.157s
    sys	0m0.034s    

Or Clojure:

    $ time clj -M fact.clj
    933262154439441526816992388562667004907159682643816214685929638
    952175999932299156089414639761565182862536979208272237582511852
    10916864000000000000000000000000N

    real	0m1.729s
    user	0m2.210s
    sys	0m0.421s

## Example REPL Session

These were copied directly from the unit test output; `eval_test.go` has more examples.

    $ l1
    > t
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
    > (eq (quote foo) (quote foo))
    t
    > (eq (quote foo) (quote bar))
    ()
    > (eq (quote foo) (quote (foo bar)))
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
    > (def fact (lambda (n) (cond ((eq 0 n) 1) (t (* n (fact (- n 1)))))))
    <lambda(n)>
    > (fact 50)
    30414093201713378043612608166064768844377641568960512000000000000
    > (len (split (fact 1000)))
    2568
    > (def fib (lambda (n) (cond ((eq 0 n) 0) ((eq 1 n) 1) (t (+ (fib (- n 1)) (fib (- n 2)))))))
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
    > ^D

    $


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
