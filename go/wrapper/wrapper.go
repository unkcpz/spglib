package wrapper

//#cgo CFLAGS: -g -Wall
//#cgo LDFLAGS: -lm ${SRCDIR}/../libs/libsymspg.a
//#include "wrapper.h"
import "C"
import "unsafe"
// import "fmt"

func DelaunayReduce(lattice []float64, symprec float64) []float64 {
  ret := C.spgo_delaunay_reduce(
    (*C.double)(unsafe.Pointer(&lattice[0])),
    (C.double)(symprec),
  )
  if ret == 0 {
    panic("spg_delaunay_reduce failed")
  }
  return lattice
}

func StandardizeCell(
  lattice []float64,
  position []float64,
  types []int,
  num_atom int,
  to_primitive int,
  no_idealize int,
  symprec float64) ([]float64, []float64, []int) {

  tp32 := to32bit(types)

  originL := make([]float64, len(lattice))
  copy(originL, lattice)
  originP := make([]float64, len(position))
  copy(originP, position)
  originT := make([]int32, len(tp32))
  copy(originT, tp32)

  n := C.spgo_standardize_cell(
    (*C.double)(unsafe.Pointer(&lattice[0])),
    (*C.double)(unsafe.Pointer(&position[0])),
    (*C.int)(unsafe.Pointer(&tp32[0])),
    (C.int)(num_atom),
    (C.int)(to_primitive),
    (C.int)(no_idealize),
    (C.double)(symprec),
  )
  if n == 0 {
    panic("spg_standardize_cell failed")
  }
  if int(n) > cap(position)/3 {
    outP := make([]float64, n*3, n*3)
    for i:=0; i<len(originP); i++ {
      outP[i] = originP[i]
    }
    outT := make([]int32, n, n)
    for i:=0; i<len(originT); i++ {
      outT[i] = originT[i]
    }
    C.spgo_standardize_cell(
      (*C.double)(unsafe.Pointer(&originL[0])),
      (*C.double)(unsafe.Pointer(&outP[0])),
      (*C.int)(unsafe.Pointer(&outT[0])),
      (C.int)(num_atom),
      (C.int)(to_primitive),
      (C.int)(no_idealize),
      (C.double)(symprec),
    )
    tp := tobit(outT)
    return originL, outP, tp
  }

  tp := tobit(tp32)
  return originL, originP, tp
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

func tobit(a []int32) []int {
	b := make([]int, len(a))
	for i := range b {
		b[i] = int(a[i])
	}
	return b
}
