package noise

import (
    "math"

    "volumetric-cloud/interpolation"
)

type Noise struct {
    PermutationTable []int // size 512

    Freq float64
    FreqFactor float64
    Amplitude float64
    AmplitudeFactor float64
    Octaves int // number of octaves
}

type ValueNoise struct {
    Table []float64 // uniform distribution (256)
    N Noise
}

type PerlinNoise struct {
    N Noise
}

func InitPerlinNoise(freq, freqFactor, amplitude, amplitudeFactor float64, octaves int) PerlinNoise {
    
    // init permutation
    permutationTable := InitPermutationTable(256)
    n := Noise{
        PermutationTable: permutationTable,
        Freq: freq,
        FreqFactor: freqFactor,
        Amplitude: amplitude,
        AmplitudeFactor: amplitudeFactor,
        Octaves: octaves,
    }

    return PerlinNoise{
        N: n,
    }
}

func (p PerlinNoise) EvalPerlinNoise(x, y, z float64) float64 {
    // Get hash coorinates (with % 256)
    xi0 := int(math.Floor(x)) & 255
    yi0 := int(math.Floor(y)) & 255
    zi0 := int(math.Floor(z)) & 255

    // upper hash coordinate
    xi1 := (xi0 + 1) & 255
    yi1 := (yi0 + 1) & 255
    zi1 := (zi0 + 1) & 255

    // Get float value of x,y,z


    xf := x - math.Floor(x)
    yf := y - math.Floor(y)
    zf := z - math.Floor(z)

    // Smooth for interpolation
    u := Fade(xf)
    v := Fade(yf)
    w := Fade(zf)

    // Get corners with hash function (permutation table)
    p1 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi0] + yi0] + zi0] // (0,0,0)
    p2 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi0] + yi1] + zi0] // (0,1,0)
    p3 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi0] + yi0] + zi1] // (0,0,1)
    p4 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi0] + yi1] + zi1] // (0,1,1)
    p5 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi1] + yi0] + zi0] // (1,0,0)
    p6 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi1] + yi1] + zi0] // (1,1,0)
    p7 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi1] + yi0] + zi1] // (1,0,1)
    p8 := p.N.PermutationTable[p.N.PermutationTable[p.N.PermutationTable[xi1] + yi1] + zi1] // (1,1,1)

    // Compute dot product
    dot1 := GradDotProduct(p1, xf, yf, zf)
    dot2 := GradDotProduct(p2, xf, yf - 1, zf)
    dot3 := GradDotProduct(p3, xf, yf, zf - 1)
    dot4 := GradDotProduct(p4, xf, yf - 1, zf - 1)
    dot5 := GradDotProduct(p5, xf - 1, yf, zf)
    dot6 := GradDotProduct(p6, xf - 1, yf - 1, zf)
    dot7 := GradDotProduct(p7, xf - 1, yf, zf - 1)
    dot8 := GradDotProduct(p8, xf - 1, yf - 1, zf - 1)

    // Linear interpolation

    // x axis interpolation
    interpP1P5 := interpolation.Lerp(u, dot1, dot5)
    interpP2P6 := interpolation.Lerp(u, dot2, dot6)
    interpP3P7 := interpolation.Lerp(u, dot3, dot7)
    interpP4P8 := interpolation.Lerp(u, dot4, dot8)

    // y axis interpolation
    interpY1 := interpolation.Lerp(v, interpP1P5, interpP2P6)
    interpY2 := interpolation.Lerp(v, interpP3P7, interpP4P8)

    // z axis
    return interpolation.Lerp(w, interpY1, interpY2)
}

func (p PerlinNoise) GeneratePerlinNoise(x, y, z float64) float64 {
    var res float64 = 0;
    amplitude := p.N.Amplitude
    freq := p.N.Freq

    for i := 0; i < p.N.Octaves; i += 1 {
        res += p.EvalPerlinNoise(x * freq, y * freq, z * freq) * amplitude

        amplitude *= p.N.AmplitudeFactor
        freq *= p.N.FreqFactor
    }
    return res
}

func InitValueNoise(freq float64) ValueNoise {
    return ValueNoise{}
}
