package main

import(
  "fmt"

  "github.com/unkcpz/spglib/go/wrapper"
)

func main() {
  latt := [][]float64{
     {15.9359925014463339,   0.0000000000000000,   0.0000000000000000},
     {-7.9679962507231625,  13.8009743407708498,   0.0000000000000000},
      {0.0000000000000000,   0.0000000000000000,   3.8919981686514271},
  }
  i := wrapper.NiggliReduce(latt, 1e-2)
  fmt.Println(i, latt)
}
