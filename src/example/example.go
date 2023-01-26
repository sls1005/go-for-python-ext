package example

//#cgo CFLAGS: -I/usr/include/python3.11
//#cgo LDFLAGS: -lpython3.11
import "C"

//export foo
func foo(x C.int) C.int {
  return x + 1
}
