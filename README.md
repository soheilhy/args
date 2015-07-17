# Introduction ![Travis Build Status](https://api.travis-ci.org/soheilhy/args.svg?branch=master "Travis Build Status")
args is a generic library for optional arguments. It is
inspired by Dave Chaney's
[functional options idea](http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis).
It can also serve the purpose of Python "kwargs" for Go programs.

# Usage
Optional arguments are defined using `New` and its typed variants:
```go
var RoundTripper = args.New()
var Timeout = args.NewDuration()
```

These arguments are basically functions that return argument values
of type `args.V`. To use these argument values the function receives
a variadic list of `args.V` and then get the value of each argument
from those values:
```go
func MyServer(args ...args.V) {
	rt := RoundTripper.Get(args)
	to := Timeout.Get(args)
	...
}
```
To use typed arguments, instead of the generic `args.V`,
you need to write a few lines of boiler-plates:

```go
var roundTripper = args.New()
var timeout = args.NewDuration()

type ServerOpt args.V
func RoundTripper(i interface{}) ServerOpt {
	return ServerOpt(roundTripper(i))
}
func Timeout(d time.Duration) ServerOpt {
	return ServerOpt(d)
}
func MyServer(opts ...ServerOpt) {
	rt := roundTripper.Get(opts)
	to := timeout.Get(opts)
	...
}
```

Note that, args is focused on easy-to-use APIs. It is not efficient
and is wasteful if the function is frequently invoked.
