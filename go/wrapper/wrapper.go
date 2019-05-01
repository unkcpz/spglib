package wrapper

//#cgo CFLAGS: -g -Wall
//#cgo LDFLAGS: -lm ${SRCDIR}/../libs/libsymspg.a
//#include "wrapper.h"
import "C"
import "unsafe"
import "fmt"

// func DelaunayReduce(lattice [][3]float64, symprec float64) int {
//   ret := C.spg_delaunay_reduce(
//     (*[3]C.double)(unsafe.Pointer(&lattice[0][0])),
//     (C.double)(symprec),
//   )
//   if ret == 0 {
//     panic("spg_delaunay_reduce failed")
//   }
//   return int(ret)
// }
//
// func StandardizeCell(
//   lattice [][3]float64,
//   position [][3]float64,
//   types []int,
//   num_atom int,
//   to_primitive int,
//   no_idealize int,
//   symprec float64) int {
//   ret := C.spg_standardize_cell(
//     (*[3]C.double)(unsafe.Pointer(&lattice[0][0])),
//     (*[3]C.double)(unsafe.Pointer(&position[0][0])),
//     (*C.int)(unsafe.Pointer(&types[0])),
//     (C.int)(num_atom),
//     (C.int)(to_primitive),
//     (C.int)(no_idealize),
//     (C.double)(symprec),
//   )
//   if ret == 0 {
//     panic("spg_standardize_cell failed")
//   }
//   return int(ret)
// }

func FindPrimitive(
  lattice []float64,
  position []float64,
  types []int,
  num_atom int,
  symprec float64) int {

  tp := to32bit(types)
  ret := C.spgo_find_primitive(
    (*C.double)(unsafe.Pointer(&lattice[0])),
    (*C.double)(unsafe.Pointer(&position[0])),
    (*C.int)(unsafe.Pointer(&tp[0])),
    (C.int)(num_atom),
    (C.double)(symprec),
  )
  fmt.Println(lattice)
  if ret == 0 {
    panic("spg_standardize_cell failed")
  }
  return int(ret)
}

// C.int is 32 bit even on a 64bit system, but Go int is 32 or 64 bit.
// So we need to convert in order to pass C int arrays.
func to32bit(a []int) []int32 {
	b := make([]int32, len(a))
	for i := range b {
		b[i] = int32(a[i])
	}
	return b
}
