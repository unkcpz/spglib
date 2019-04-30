package wrapper

//#cgo CFLAGS: -g -Wall
//#cgo LDFLAGS: -lm ${SRCDIR}/../libs/libsymspg.a
//#include "spglib.h"
import "C"
import "unsafe"
import "fmt"

func DelaunayReduce(lattice [][3]float64, symprec float64) int {
  ret := C.spg_delaunay_reduce(
    (*[3]C.double)(unsafe.Pointer(&lattice[0][0])),
    (C.double)(symprec),
  )
  if ret == 0 {
    panic("spg_delaunay_reduce failed")
  }
  return int(ret)
}

func StandardizeCell(
  lattice [][3]float64,
  position [][3]float64,
  types []int,
  num_atom int,
  to_primitive int,
  no_idealize int,
  symprec float64) int {
  ret := C.spg_standardize_cell(
    (*[3]C.double)(unsafe.Pointer(&lattice[0][0])),
    (*[3]C.double)(unsafe.Pointer(&position[0][0])),
    (*C.int)(unsafe.Pointer(&types[0])),
    (C.int)(num_atom),
    (C.int)(to_primitive),
    (C.int)(no_idealize),
    (C.double)(symprec),
  )
  if ret == 0 {
    panic("spg_standardize_cell failed")
  }
  return int(ret)
}

func FindPrimitive(
  lattice [][3]float64,
  position [][3]float64,
  types []int,
  num_atom int,
  symprec float64) int {
  n := 0
  latt := make([]float64, 9, 9)
  fmt.Println(latt)
  for _, row := range lattice {
    for _, v := range row {
      latt[n] = v
      n++
    }
  }
  n = 0
  pos := make([]float64, 3*num_atom, 3*num_atom)
  fmt.Println(pos)
  for _, row := range position {
    for _, v := range row {
      pos[n] = v
      n++
    }
  }
  fmt.Println(pos)
  ret := C.spg_find_primitive(
    (*[3]C.double)(unsafe.Pointer(&latt[0])),
    (*[3]C.double)(unsafe.Pointer(&pos[0])),
    (*C.int)(unsafe.Pointer(&types[0])),
    (C.int)(num_atom),
    (C.double)(symprec),
  )
  fmt.Println(latt)
  if ret == 0 {
    panic("spg_standardize_cell failed")
  }
  return int(ret)
}
