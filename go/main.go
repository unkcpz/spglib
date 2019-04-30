package main

import(
  "github.com/unkcpz/spglib/go/wrapper"
)

func main() {
  latt = [][]float64{
    {2.456, 0.0, 0.0},
    {11.052, 2.127, 0.0},
    {0.0, 0.0, 14.0},
  }
  wrapper.NiggliReduce(latt, 1e-5)
  fmt.Println(latt)
}
