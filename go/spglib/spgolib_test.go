package spglib

import (
  // "fmt"
  "testing"
)

func TestNewDataset(t *testing.T) {
  ds := NewDataset(
    []float64{4,0,0,0,4,0,0,0,4},
    []float64{0,0,0,0.5,0.5,0.5},
    []int{1,1},
    1e-5)

  got_0 := ds.SpaceNumber
  if got_0 != 229 {
    t.Errorf("spacegroupNumber, expect 229, got %d", got_0)
  }

  got_1 := ds.SpaceSymbol
  if got_1 != "Im-3m" {
    t.Errorf("international_symbol, expect Im-3m, got %s", got_1)
  }

  got_2 := ds.HallNumber
  if got_2 != 529 {
    t.Errorf("hallNumber, expect 229, got %d", got_2)
  }

  got_3 := ds.HallSymbol
  if got_3 != "-I 4 2 3" {
    t.Errorf("international_symbol, expect '-I 4 2 3', got %s", got_3)
  }

  got_4 := ds.Nops
  if got_4 != 96 {
    t.Errorf("number of operations, expect 96, got %d", got_4)
  }
  // fmt.Println(ds.Rotations)
  // fmt.Println(ds.Translations)
}
