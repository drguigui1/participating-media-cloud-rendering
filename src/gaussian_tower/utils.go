package gaussian_tower

func Transpose3(mat []float64) []float64{
	return []float64 {mat[0], mat[3], mat[6], mat[1], mat[4], mat[7], mat[2], mat[5], mat[8]}
}

func Multiplication3(mat1 []float64, mat2 []float64) []float64{
	return []float64 {
		mat1[0] * mat2[0] + mat1[1] * mat2[3] + mat1[2] * mat2[6],
		mat1[0] * mat2[1] + mat1[1] * mat2[4] + mat1[2] * mat2[7],
		mat1[0] * mat2[2] + mat1[1] * mat2[5] + mat1[2] * mat2[8],
		mat1[3] * mat2[0] + mat1[4] * mat2[3] + mat1[5] * mat2[6],
		mat1[3] * mat2[1] + mat1[4] * mat2[4] + mat1[5] * mat2[7],
		mat1[3] * mat2[2] + mat1[4] * mat2[5] + mat1[5] * mat2[8],
		mat1[6] * mat2[0] + mat1[7] * mat2[3] + mat1[8] * mat2[6],
		mat1[6] * mat2[1] + mat1[7] * mat2[4] + mat1[8] * mat2[7],
		mat1[6] * mat2[2] + mat1[7] * mat2[5] + mat1[8] * mat2[8],
	}
}