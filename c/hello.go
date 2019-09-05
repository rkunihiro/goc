package c

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lhello
// #include "hello.h"
import "C"

func Hello() {
	C.hello()
}
