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
int spgo_find_primitive(double *lattice,
                         double *position,
                         int *types,
                         const int num_atom,
                         const double symprec);
// static void showgo_cell(double lattice[],
// 		      double position[],
// 		      const int types[],
// 		      const int num_atom);
// static void show_cell(double lattice[3][3],
// 		      double position[][3],
// 		      const int types[],
// 		      const int num_atom);

#endif
