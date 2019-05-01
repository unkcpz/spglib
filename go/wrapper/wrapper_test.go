package wrapper

import (
  "testing"
)

func TestDelaunayReduce(t *testing.T) {
  lattice := []float64{4, 0, 0, 2, 4, 0, 0, 0, 4}
  origin := make([]float64, len(lattice))
  copy(origin, lattice)

  lattice = DelaunayReduce(lattice, 1e-5)

  expected := []float64{0, 0, -4, 0, -4, -2, -4, 0, 0}
  for i, _ := range lattice {
    if lattice[i] - expected[i] > 1e-5 {
      t.Errorf("delaunay reduce lattice: %v, expected %v, got %v.", origin, expected, lattice)
      break
    }
  }
}

func TestStandardizeCellBCCRefine(t *testing.T) {
  // BCC to its Conventional
  lattice := []float64{
    -2, 2, 2,
    2, -2, 2,
    2, 2, -2,
  }
  originL := make([]float64, len(lattice))
  copy(originL, lattice)

  position := []float64{
    0, 0, 0,
  }
  originP := make([]float64, len(position))
  copy(originP, position)

  types := []int{1}
  n_atoms := 1
  lattice, position, types = StandardizeCell(lattice, position, types, n_atoms, 0, 0, 1e-5)

  expectedL := []float64{4, 0, 0, 0, 4, 0, 0, 0, 4}
  expectedP := []float64{0, 0, 0, 0.5, 0.5, 0.5}
  expectedT := []int{1, 1}
  for i, _ := range expectedL {
    if lattice[i] - expectedL[i] > 1e-5 {
      t.Error("StandardizeCell error")
      break
    }
  }
  for i, _ := range expectedP{
    if position[i] - expectedP[i] > 1e-5 {
      t.Error("StandardizeCell error")
      break
    }
  }
  for i, _ := range expectedT {
    if types[i] != expectedT[i] {
      t.Error("StandardizeCell error")
      break
    }
  }
}

func TestStandardizeCellBCCPrimitive(t *testing.T) {
  // BCC to its primitive
  lattice := []float64{
    4, 0, 0,
    0, 4, 0,
    0, 0, 4,
  }
  originL := make([]float64, len(lattice))
  copy(originL, lattice)

  position := []float64{
    0, 0, 0,
    0.5, 0.5, 0.5,
  }
  originP := make([]float64, len(position))
  copy(originP, position)

  types := []int{1, 1}
  n_atoms := 2
  lattice, position, types = StandardizeCell(lattice, position, types, n_atoms, 1, 1, 1e-5)

  expectedL := []float64{-2, 2, 2, 2, -2, 2, 2, 2, -2}
  expectedP := []float64{0, 0, 0}
  expectedT := []int{1}
  for i, _ := range expectedL {
    if lattice[i] - expectedL[i] > 1e-5 {
      t.Errorf("got %v, expected %v", lattice, expectedL)
      break
    }
  }
  for i, _ := range expectedP{
    if position[i] - expectedP[i] > 1e-5 {
      t.Error("StandardizeCell error")
      break
    }
  }
  for i, _ := range expectedT {
    if types[i] != expectedT[i] {
      t.Error("StandardizeCell error")
      break
    }
  }
}
