#include "spglib.h"
#include "wrapper.h"

int spgo_delaunay_reduce(double lattice[], const double symprec) {
  double latt[3][3];
  mat_flat_3D(lattice, latt, 3);
  int r = spg_delaunay_reduce(latt, symprec);
  flat_mat_3D(latt, lattice);
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
  int r = spg_standardize_cell(latt, pos, types, num_atom, to_primitive, no_idealize, symprec);
  flat_mat_3D(latt, lattice, 3);
  flat_mat_3D(pos, position, num_atom);
	return r;
}

int spgo_find_primitive(double lattice[],
                        double position[],
                        int types[],
                        const int num_atom,
                        const double symprec) {

	double latt[3][3];
	double pos[num_atom][3];
  mat_flat_3D(lattice, latt, 3);
  mat_flat_3D(position, pos, num_atom);
  int r = spg_find_primitive(latt, pos, types, num_atom, symprec);
  flat_mat_3D(latt, lattice, 3);
  flat_mat_3D(pos, position, num_atom);
	return r;
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
