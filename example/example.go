package example

//#cgo CFLAGS: -I/usr/include/python3.10
//#cgo LDFLAGS: -lpython3.10
import "C"

//export foo
func foo(x C.int) C.int {
  return x + 1
}
