package vector3

import (
    "math"
    "fmt"
)

type Vector3 struct {
    X float64
    Y float64
    Z float64
}

func InitVector3(x, y, z float64) Vector3 {
    return Vector3{
        X: x,
        Y: y,
        Z: z,
    }
}

func (v *Vector3) Display() {
    fmt.Println("Vector3")
    fmt.Println("x: ", v.X)
    fmt.Println("y: ", v.Y)
    fmt.Println("z: ", v.Z)
    fmt.Println("")
}

func (v *Vector3) Add(val float64) {
    v.X += val
    v.Y += val
    v.Z += val
}

func (v *Vector3) AddVector3(other Vector3) {
    v.X += other.X
    v.Y += other.Y
    v.Z += other.Z
}

func (v *Vector3) Sub(val float64) {
    v.X -= val
    v.Y -= val
    v.Z -= val
}

func (v *Vector3) Neg() {
    v.X = -v.X
    v.Y = -v.Y
    v.Z = -v.Z
}

func (v *Vector3) OneMinus() {
    v.X = 1 - v.X
    v.Y = 1 - v.Y
    v.Z = 1 - v.Z
}

func (v *Vector3) Mul(val float64) {
    v.X *= val
    v.Y *= val
    v.Z *= val
}

func (v *Vector3) Div(val float64) {
    v.X /= val
    v.Y /= val
    v.Z /= val
}

func (v *Vector3) Floor() {
    v.X = math.Floor(v.X)
    v.Y = math.Floor(v.Y)
    v.Z = math.Floor(v.Z)
}

func (v *Vector3) Clamp(low, high float64) {
    if v.X < low {
        v.X = low
    } else if v.X > high {
        v.X = high
    }

    if v.Y < low {
        v.Y = low
    } else if v.Y > high {
        v.Y = high
    }

    if v.Z < low {
        v.Z = low
    } else if v.Z > high {
        v.Z = high
    }
}

func (v *Vector3) Copy() Vector3 {
    return InitVector3(v.X, v.Y, v.Z)
}

func (v Vector3) Length() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func AddVector3(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z)
}

func SubVector3(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z)
}

func HadamarProduct(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z)
}

func MulVector3Scalar(v Vector3, val float64) Vector3 {
    return Vector3{
        X: v.X * val,
        Y: v.Y * val,
        Z: v.Z * val,
    }
}

func DivVector3(v Vector3, val float64) Vector3 {
    return Vector3{
        X: v.X / val,
        Y: v.Y / val,
        Z: v.Z / val,
    }
}

func NegVector3(v Vector3) Vector3 {
    return InitVector3(-v.X, -v.Y, -v.Z)
}

func UnitVector(v Vector3) Vector3 {
    return DivVector3(v, v.Length())
}

func DotProduct(v1, v2 Vector3) float64 {
    return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func CrossProduct(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.Y * v2.Z - v1.Z * v2.Y,
                       v1.Z * v2.X - v1.X * v2.Z,
                       v1.X * v2.Y - v1.Y * v2.X)
}

// m: size 9 (3x3)
func MultMatVec3(m []float64, v Vector3) Vector3 {
    return InitVector3(
        m[0] * v.X + m[1] * v.Y + m[2] * v.Z,
        m[3] * v.X + m[4] * v.Y + m[5] * v.Z,
        m[6] * v.X + m[7] * v.Y + m[8] * v.Z,
    )
}
