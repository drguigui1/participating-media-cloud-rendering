package noise

type ValueNoise struct {
    Table []float64 // uniform distribution (256)
    PermutationTable []int // (512)

    Freq float64
    FreqFactor float64
    Amplitude float64
    AmplitudeFactor float64
    Octaves int // number of octaves
}

func InitValueNoise(freq float64) ValueNoise {
    return ValueNoise{}
}
