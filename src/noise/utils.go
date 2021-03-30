package noise

import (
    "math/rand"
)

func Arange(size int) []int {
    slice := make([]int, size)
    for i := 0; i < size; i += 1 {
        slice[i] = i
    }

    return slice
}

func Shuffle(s *[]int) {
    size := len((*s))
    for i := 0; i < size; i += 1 {
        // random value
        rdPos := rand.Intn(size)
        tmp := (*s)[i]
        (*s)[i] = (*s)[rdPos]
        (*s)[rdPos] = tmp
    }
}

func Stack(s1 []int, s2[]int) []int {
    res := make([]int, len(s1) + len(s2))

    for i := 0; i < len(s1); i += 1 {
        res[i] = s1[i]
    }

    for i := 0; i < len(s1); i += 1 {
        res[i + len(s1)] = s2[i]
    }

    return res
}

func InitPermutationTable(size int) []int {
    permutationTable := Arange(size)
    Shuffle(&permutationTable)
    return Stack(permutationTable, permutationTable)
}

// Dot product between x,y,z and the gradient vectors
// (according to the value in the hash table)
func GradDotProduct(hashRes int, x, y, z float64) float64 {
    m := []float64{
        x + y,
        -x + y,
        x - y,
        -x - y,
        x + z,
        -x + z,
        x - z,
        -x - z,
        y + z,
        -y + z,
        y - z,
        -y - z,
        x + y,
        -y + z,
        -x + y,
        -y - z,
    }

    return m[hashRes & 15] // % 16
}
