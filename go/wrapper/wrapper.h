#ifndef _WRAPPER_H
#define _WRAPPER_H

int spgo_delaunay_reduce(double lattice[], const double symprec);
int spgo_standardize_cell(double lattice[],
                        double position[],
                        int types[],
                        const int num_atom,
                        const int to_primitive,
                        const int no_idealize,
                        const double symprec);
// int spgo_find_primitive(double *lattice,
//                          double *position,
//                          int *types,
//                          const int num_atom,
//                          const double symprec);
void flat_mat_3D(double mat[][3], double flat[], int n);
void mat_flat_3D(double flat[], double mat[][3], int n);

#endif
