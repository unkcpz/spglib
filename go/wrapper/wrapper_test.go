package wrapper

import (
  "testing"
)

func TestDelaunayReduce(t *testing.T) {
  lattice := []float64{4, 0, 0, 2, 4, 0, 0, 0, 4}
  origin := make([]float64, len(lattice))
  copy(origin, lattice)
  r := DelaunayReduce(lattice, 1e-5)
  if r != 1 {
    t.Errorf("delaunay reduce lattice: %v, failed.", origin)
  }
  expected := []float64{0, 0, -4, 0, -4, -2, -4, 0, 0}
  for i, _ := range lattice {
    if lattice[i] - expected[i] > 1e-5 {
      t.Errorf("delaunay reduce lattice: %v, expected %v, got %v.", origin, expected, lattice)
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
  copy(originP, lattice)

  types := []int{1, 1}
  n_atoms := 2
  n := StandardizeCell(lattice, position, types, n_atoms, 1, 0, 1e-5)
  // atom number of primitive is 1
  if n != 1 {
    t.Errorf("primitive of cell: (latt)=%v, (pos)=%v, (types)=%v, n_atoms is %d, expect 1", originL, originP, []int{1, 1}, n)
  }
  expectedL := []float64{-2, 2, 2, 2, -2, 2, 2, 2, -2}
  expectedP := []float64{0, 0, 0,0,0,0}
  for i, _ := range lattice {
    if lattice[i] - expectedL[i] > 1e-5 {
      t.Errorf("expect primitive lattice %v, got %v", expectedL, lattice)
      break
    }
  }
  for i:=0; i<n; i++ {
    if position[i] - expectedP[i] > 1e-5 {
      t.Errorf("expect primitive position %v, got %v", expectedP, position)
      break
    }
  }
}

func TestStandardizeCellBCCRefine(t *testing.T) {
  // BCC to its primitive
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
  copy(originP, lattice)

  types := []int{1}
  n_atoms := 1
  n := StandardizeCell(lattice, position, types, n_atoms, 0, 0, 1e-5)
  // atom number of primitive is 1
  if n != 2 {
    t.Errorf("primitive of cell: (latt)=%v, (pos)=%v, (types)=%v, n_atoms is %d, expect 2", originL, originP, []int{1}, n)
  }
  expectedL := []float64{4, 0, 0, 0, 4, 0, 0, 0, 4}
  expectedP := []float64{0, 0, 0, 0.5, 0.5, 0.5}
  for i, _ := range lattice {
    if lattice[i] - expectedL[i] > 1e-5 {
      t.Errorf("expect primitive lattice %v, got %v", expectedL, lattice)
      break
    }
  }
  for i, _ := range position {
    if position[i] - expectedP[i] > 1e-5 {
      t.Errorf("expect primitive position %v, got %v", expectedP, position)
      break
    }
  }
}
