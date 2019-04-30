package wrapper

//#cgo CFLAGS: -g -Wall
//#cgo LDFLAGS: -lm ${SRCDIR}/../libs/libsymspg.a
//#include "spglib.h"
import "C"
import "unsafe"

func NiggliReduce(lattice [][]float64, symprec float64) int {
  ret := C.spg_niggli_reduce(
    (*[3]C.double)(unsafe.Pointer(&lattice[0])),
    (C.double)(symprec),
  )
  if ret == 0 {
    panic("spg_niggli_reduce failed")
  }
  return int(ret)
}
