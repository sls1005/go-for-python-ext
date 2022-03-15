# Go for Python extensions

This contains an example of a Python [extension](https://docs.python.org/3/extending/extending.html) written in [Go](https://github.com/golang/go) and [Cython](https://github.com/cython/cython).

It is tested on Linux. Some of the files need to be changed in order to compile on Windows.

### Build

To build the example, you need...

+ Python 2.7+ or 3.3+ (CPython)
+ Cython 0.20+
+ Go 1.12+


From the top directory:

```sh
$ make
```

### How it works

When `import "C"` is used in a `.go` file, all the `.c` files in the same package are also compiled and linked together. We can use Cython to generate the C source necessary for a Python extension and then ask the Go compiler to compile and link everything.

### References

<https://pkg.go.dev/cmd/cgo>
