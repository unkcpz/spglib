// Internally used package spglib provides low-level wrappers for libspglib
package spglib

//#cgo CFLAGS: -std=gnu99 -g -Wall
//#cgo LDFLAGS: -lm
//#inclue "spglib.h"
import "C"
import "unsafe"

// Wrapper for spg_niggli_reduce.
// Internal but exported for use by package gospglib.
func NiggliReduce(lattice [][]float32, symprec float32) int32 {
  // check lattice dimension here
  ret := C.spg_niggli_reduce(
    (**C.float)(unsafe.Pointer(&lattice[0][0])),
    (*C.int)(symprec)
  )

  if ret == 0 {
    panic("spg_niggli_reduce failed")
  }
}

// func CDoubleMatrixPtr(matrix [][]float32) *float32{
//   var p *float32
//   if len(matrix) > 0 && len(matrix[0]) > 0 {
//     p = (*float32)(unsafe.Pointer(&matrix[0][0]))
//   }
// }
