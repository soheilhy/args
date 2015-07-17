# Introduction ![Travis Build Status](https://api.travis-ci.org/soheilhy/args.svg?branch=master "Travis Build Status")
args is a generic library for optional arguments. It is
inspired by Dave Chaney's
[functional options idea](http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis).
It can also serve the purpose of Python "kwargs" for Go programs.

# Usage
Optional arguments are defined using `New` and its typed variants:
```go
var Port = args.NewInt()
var RoundTripper = args.New(Default(http.DefaultTransport))
var Timeout = args.NewDuration(Flag("timeout", 10*time.Second, "timeout"))
```

Note that `args` can retrieve the default value from flags as well.

These arguments are basically functions that return argument values
of type `args.V`. To use these argument values the function receives
a variadic list of `args.V` and then get the value of each argument
from those values:
```go
func MyServer(args ...args.V) {
	port := Port.Get(args)
	rt := RoundTripper.Get(args)
	to := Timeout.Get(args)
	...
}
```

The client calls your API passing some of these arguments:
```go
MyServer()
MyServer(Timeout(1 * time.Second))
MyServer(Timeout(2 * time.Second), RoundTripper(MyTransport))
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

# API
* [Go Doc](https://godoc.org/github.com/soheilhy/args)
