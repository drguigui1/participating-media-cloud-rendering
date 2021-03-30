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
