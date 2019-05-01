#include <stdio.h>
#include "spglib.h"
#include "wrapper.h"

int spgo_delaunay_reduce(double lattice[], const double symprec) {
  double latt[3][3];
  mat_flat_3D(lattice, latt, 3);
  int r = spg_delaunay_reduce(latt, symprec);
  flat_mat_3D(latt, lattice, 3);
  return r;
}

int spgo_standardize_cell(double lattice[],
                        double position[],
                        int types[],
                        const int num_atom,
                        const int to_primitive,
                        const int no_idealize,
                        const double symprec) {

	double latt[3][3];
	double pos[num_atom][3];
  mat_flat_3D(lattice, latt, 3);
  mat_flat_3D(position, pos, num_atom);
  int n = spg_standardize_cell(latt, pos, types, num_atom, to_primitive, no_idealize, symprec);
  flat_mat_3D(latt, lattice, 3);
  flat_mat_3D(pos, position, n);
	return n;
}

SpglibDataset * spgo_get_dataset(double lattice[],
                                 double position[],
                                 int types[],
                                 const int num_atom,
                                 const double symprec) {

	double latt[3][3];
	double pos[num_atom][3];
  SpglibDataset *dataset;

  mat_flat_3D(lattice, latt, 3);
  mat_flat_3D(position, pos, num_atom);
  dataset = spg_get_dataset(latt, pos, types, num_atom, symprec);
  return dataset;
}

void spgo_free_dataset(SpglibDataset *dataset) {
  spg_free_dataset(dataset);
}

int test_dataset(SpglibDataset *dataset) {
  int i, j;
  for (i = 0; i < dataset->n_operations; i++) {
    printf("--- %d ---\n", i + 1);
    for (j = 0; j < 3; j++) {
      printf("%2d %2d %2d\n",
	     dataset->rotations[i][j][0],
	     dataset->rotations[i][j][1],
	     dataset->rotations[i][j][2]);
    }
    printf("%f %f %f\n",
	   dataset->translations[i][0],
	   dataset->translations[i][1],
	   dataset->translations[i][2]);
  }
  return (int)(sizeof(dataset->translations));
}

void flat_mat_3D(double mat[][3], double flat[], int n) {
  double *pmat = &mat[0][0];
  double *pflat = &flat[0];
  for(int i=0; i<n*3; i++) {
    *pflat++ = *pmat++;
  }
}

void mat_flat_3D(double flat[], double mat[][3], int n) {
  double *pmat = &mat[0][0];
  double *pflat = &flat[0];
  for(int i=0; i<n*3; i++) {
    *pmat++ = *pflat++;
  }
}
