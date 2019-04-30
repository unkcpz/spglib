package main

import(
  "fmt"

  "github.com/unkcpz/spglib/go/wrapper"
)

func main() {
  latt := [3][3]float64{
    {0.0, 2.0, 0.0},
    {4.0, 0.0, 0.0},
    {0.0, 0.0, 14.0},
  }
  i := wrapper.DelaunayReduce(latt, 1e-2)
  fmt.Println(i, latt)
}
