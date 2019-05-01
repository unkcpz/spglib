#include <stdio.h>
#include "spglib.h"
#include "wrapper.h"


int spgo_find_primitive(double lattice[],
                        double position[],
                        int types[],
                        const int num_atom,
                        const double symprec) {

	double latt[3][3];
	double pos[num_atom][3];
	for(int i=0; i<3; i++){
		for(int j=0; j<3; j++){
			latt[i][j] = lattice[3*i+j];
		}
	}
	for(int i=0; i<num_atom; i++){
		for(int j=0; j<3; j++){
			pos[i][j] = position[3*i+j];
		}
	}
  showgo_cell(lattice, position, types, num_atom);
	int r=0;
  r = spg_find_primitive(latt, pos, types, num_atom, symprec);
	for(int i=0; i<3; i++){
		for(int j=0; j<3; j++){
			lattice[3*i+j] = latt[i][j];
		}
	}
	for(int i=0; i<num_atom; i++){
		for(int j=0; j<3; j++){
			position[3*i+j] = pos[i][j];
		}
	}
  showgo_cell(lattice, position, types, num_atom);

	return r;
}

static void showgo_cell(double lattice[],
		      double position[],
		      const int types[],
		      const int num_atom)
{
  int i;
	double latt[3][3];
	double pos[num_atom][3];
	for(int i=0; i<3; i++){
		for(int j=0; j<3; j++){
			latt[i][j] = lattice[3*i+j];
		}
	}
	for(int i=0; i<num_atom; i++){
		for(int j=0; j<3; j++){
			pos[i][j] = position[3*i+j];
		}
	}

  printf("Lattice parameter:\n");
  for (i = 0; i < 3; i++) {
    printf("%f %f %f\n", latt[0][i], latt[1][i], latt[2][i]);
  }
  printf("Atomic positions:\n");
  for (i = 0; i < num_atom; i++) {
    printf("%d: %f %f %f\n",
	   types[i], pos[i][0], pos[i][1], pos[i][2]);
  }
}
