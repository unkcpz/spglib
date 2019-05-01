#ifndef _WRAPPER_H
#define _WRAPPER_H

int spgo_find_primitive(double *lattice,
                         double *position,
                         int *types,
                         const int num_atom,
                         const double symprec);
static void showgo_cell(double lattice[],
		      double position[],
		      const int types[],
		      const int num_atom);

#endif
