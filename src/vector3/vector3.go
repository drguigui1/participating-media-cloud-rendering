package vector3

import (
    "math"
)

type Vector3 struct {
    x float64
    y float64
    z float64
}

func InitVector3(x, y, z float64) Vector3 {
    return Vector3{
        x: x,
        y: y,
        z: z,
    }
}

func (v *Vector3) Add(val float64) {
    v.x += val
    v.y += val
    v.z += val
}

func (v *Vector3) Sub(val float64) {
    v.x -= val
    v.y -= val
    v.z -= val
}

func (v *Vector3) Neg() {
    v.x = -v.x
    v.y = -v.y
    v.z = -v.z
}

func (v *Vector3) Mul(val float64) {
    v.x *= val
    v.y *= val
    v.z *= val
}

func (v *Vector3) Div(val float64) {
    // TODO (Check if val == 0)
    v.x /= val
    v.y /= val
    v.z /= val
}

func (v Vector3) Length() float64 {
    return math.Sqrt(v.x * v.x + v.y * v.y + v.z * v.z)
}

func AddVector3(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.x + v2.x, v1.y + v2.y, v1.z + v2.z)
}

func SubVector3(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.x - v2.x, v1.y - v2.y, v1.z - v2.z)
}

func MulVector3(v Vector3, val float64) Vector3 {
    return Vector3{
        x: v.x * val,
        y: v.y * val,
        z: v.z * val,
    }
}

func DivVector3(v Vector3, val float64) Vector3 {
    return Vector3{
        x: v.x / val,
        y: v.y / val,
        z: v.z / val,
    }
}

func UnitVector(v Vector3) Vector3 {
    return DivVector3(v, v.Length())
}

func DotProduct(v1, v2 Vector3) float64 {
    return v1.x * v2.x + v1.y * v2.y + v1.z * v2.z
}

func CrossProduct(v1, v2 Vector3) Vector3 {
    return InitVector3(v1.y * v2.z - v1.z * v2.y,
                       v1.z * v2.x - v1.x * v2.z,
                       v1.x * v2.y - v1.y * v2.x)
}

// m: size 9 (3x3)
func MultMatVec3(m []float64, v Vector3) Vector3 {
    return InitVector3(
        m[0] * v.x + m[1] * v.y + m[2] * v.z,
        m[3] * v.x + m[4] * v.y + m[5] * v.z,
        m[6] * v.x + m[7] * v.y + m[8] * v.z,
    )
}
