package main

import(
  "fmt"

  "github.com/unkcpz/spglib/go/wrapper"
)

func main() {
  latt := []float64{
    4, 0, 0,
    0, 4, 0,
    0, 0, 4,
  }
  positions := []float64 {
    0.0, 0.0, 0.0,
    0.5, 0.0, 0.5,
    0.0, 0.5, 0.5,
    0.5, 0.5, 0.0,
  }
  types := []int{1, 1, 1, 1}
  num_atom := 4
  // i := wrapper.DelaunayReduce(latt, 1e-2)
  i := wrapper.FindPrimitive(
    latt,
    positions,
    types,
    num_atom,
    1e-5,
  )
  fmt.Println(i, latt, positions)
}
