# l1

<img src="/l1.jpg" width="400">

![build](https://github.com/eigenhombre/l1/actions/workflows/build.yml/badge.svg)

First attempt at a simple Lisp in Go.

# Implemented
- Lexing and parsing
- Numbers as `big.Int`
- Eval
  - `t` as True, `()` as `nil`
  - Atoms bind to values in an environment
  - `quote`, `car`, `cdr`, `cons`, `cond`
  - Arithmetical operators `+-*/`
  - Side effects, e.g. `(def a 1234)`
- (Some) simple error handling

# Usage / Examples

Check out this repo and `cd` to it. Then,

    go test
    go build
    go install

Then, from anywhere, `l1` will start your REPL:

    $ l1
    > t
    t
    > ()
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
    > (cond (() 1) (2 3))
    3
    > 1
    1
    > -5
    -5
    > (* 12349807213490872130987 12349807213490872130987)
    152517738210391179737088822267441718485594169
    > (+)
    0
    > (+ 1 1 2 3)
    7
    > (+ 1 1)
    2
    > (eq (quote foo) (quote foo))
    t
    > (eq (quote foo) (quote bar))
    ()
    > (def a 1234567890)
    1234567890
    > (* a a a a a)
    2867971860299718107233761438093672048294900000
    > (
    unbalanced parens
    > )
    unexpected right paren
    > ^D
    $

These were copied directly from the unit test output; `eval_test.go` has more examples.

A `Makefile` exists for convenience, and a `Dockerfile` is used for a GitHub action CI build.

# Planned Features

- [Lisp 1](https://en.wikipedia.org/wiki/Common_Lisp#The_function_namespace);
- Sufficient power to [implement itself](http://www.paulgraham.com/rootsoflisp.html);
- Implement math as bignums from the get-go;
- Curses-based terminal control for text games, command line utilities, ...;

# Goals

- Improve my Go skills;
- Build a small, fast-loading Lisp that I can extend how I like;
- Learn more about [Lisp as a model for computation](http://www.paulgraham.com/rootsoflisp.html).

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
