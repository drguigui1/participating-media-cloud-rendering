package gaussian_tower

import (
	"reflect"
	"testing"
)

func TestMatMult3(t *testing.T) {
	mat1 := []float64{10, 20, 10, 4, 5, 6, 2, 3, 5}
	mat2 := []float64{3, 2, 4, 3, 3, 9, 4, 4, 2}

	res := Multiplication3(mat1, mat2)
	ref := []float64{130, 120, 240, 51, 47, 73, 35, 33, 45}
	if !reflect.DeepEqual(ref, res) {
		t.Errorf("Error 'TestMatMult3'\n")
		t.Errorf("res: %v\n", res)
		t.Errorf("ref: %v\n", ref)
	}
}

func TestMatTranspose3(t *testing.T) {
	mat1 := []float64{10, 20, 10, 4, 5, 6, 2, 3, 5}

	res := Transpose3(mat1)
	ref := []float64{10, 4, 2, 20, 5, 3, 10, 6, 5}
	if !reflect.DeepEqual(ref, res) {
		t.Errorf("Error 'TestMatTranspose3'\n")
		t.Errorf("res: %v\n", res)
		t.Errorf("ref: %v\n", ref)
	}
}