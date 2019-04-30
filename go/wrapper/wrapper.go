package wrapper

//#cgo CFLAGS: -g -Wall
//#cgo LDFLAGS: -lm ${SRCDIR}/../libs/libsymspg.a
//#include "spglib.h"
import "C"
import "unsafe"

func DelaunayReduce(lattice [3][3]float64, symprec float64) int {
  ret := C.spg_delaunay_reduce(
    (*[3]C.double)(unsafe.Pointer(&lattice[0][0])),
    (C.double)(symprec),
  )
  if ret == 0 {
    panic("spg_delaunay_reduce failed")
  }
  return int(ret)
}
