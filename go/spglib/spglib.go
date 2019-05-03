package spglib

import (
  "fmt"

  "gonum.org/v1/gonum/mat"
  "github.com/unkcpz/spglib/go/wrapper"
)

type Cell struct {
  // Value rather than ptr here for non-mutable Cell

  // Lattice
  Lattice mat.Dense
  //
  Positions mat.Dense
  //
  Types []int
  //
  Natoms int
}

func NewCell(
  lattice []float64,
  positions []float64,
  types []int) (*Cell, error) {

  // TODO: add check here for valid Cell
  if len(lattice) != 9 {
    return nil, fmt.Errorf("expext 9 value as lattice")
  }
  if len(positions) != 3*len(types) {
    return nil, fmt.Errorf("atom number not compatible, len(types)=%d, len(positions)/3=%d/3", len(types), len(positions))
  }

  n := len(types)
  cell := &Cell {
    Lattice: *mat.NewDense(3, 3, lattice),
    Positions: *mat.NewDense(n, 3, positions),
    Types: types,
    Natoms: n,
  }
  return cell, nil
}


type Dataset struct {
  SpaceNumber int
  SpaceSymbol string
  HallNumber int
  HallSymbol string

  Nops int
  Rotations []mat.Dense
  Translations []mat.VecDense
}

func NewDataset(cell Cell, eps float64) *Dataset {
  ds := wrapper.NewSpglibDataset(
    MatData(cell.Lattice),
    MatData(cell.Positions),
    cell.Types,
    cell.Natoms,
    eps,
  )

  nops := ds.Noperations
  rots := make([]mat.Dense, nops, nops)
  trans := make([]mat.VecDense, nops, nops)
  for i:=0; i<nops; i++ {
    rots[i] = *mat.NewDense(3, 3, intToFloat(ds.Rotations[i*9:i*9+9]))
    trans[i] = *mat.NewVecDense(3, ds.Translations[i*3:i*3+3])
  }
  rds := &Dataset{
    SpaceNumber: ds.SpacegroupNumber,
    SpaceSymbol: ds.InternationalSymbol,
    HallNumber: ds.HallNumber,
    HallSymbol: ds.HallSymbol,
    Nops: ds.Noperations,
    Rotations: rots,
    Translations: trans,
  }
  return rds
}

func intToFloat(a []int) []float64 {
  b := make([]float64, len(a))
  for i := range b {
    b[i] = float64(a[i])
  }
  return b
}

func MatData(mat mat.Dense) []float64 {
  blasM := mat.RawMatrix()
  return blasM.Data
}
