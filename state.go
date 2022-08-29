package main

import "fmt"

var gensymCounter = 0

// Currently we only have one thread.  If we ever have concurrency, ensure
// thread safety.
func gensym(prefix string) string {
	gensymCounter++
	return fmt.Sprintf("<gensym%s-%d>", prefix, gensymCounter)
}
