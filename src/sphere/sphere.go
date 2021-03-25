package sphere

import (
    "math"

    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
)

type Sphere struct {
    Center vector3.Vector3
    Radius float64
}

func InitSphere(center vector3.Vector3, radius float64) Sphere {
    return Sphere{
        Center: center,
        Radius: radius,
    }
}

/*
** Eq of the ray: P(x) = A + tb
** with A: origin and b: direction
**
** Eq of the sphere: (P-C)(P-C) = r^2
** with P=(x,y,z), C=center, r=radius
**
** So intersection between sphere and ray:
** t^2 * b.b + 2t*b.(A-C) + (A-C).(A-C) - r^2 = 0
*/
func (s Sphere) Hit(ray ray.Ray) (float64, vector3.Vector3, bool) {
    diff := vector3.SubVector3(ray.Origin, s.Center)

    a := vector3.DotProduct(ray.Direction, ray.Direction)
    b := vector3.DotProduct(vector3.MulVector3(ray.Direction, 2.0), diff)
    c := vector3.DotProduct(diff, diff) - (s.Radius * s.Radius)

    delta := b * b - 4.0 * a * c

    if delta < 0 {
        return 0.0, vector3.Vector3{}, false
    }

    root := ((-b - math.Sqrt(delta)) / (2.0 * a))
    if root < 0 {
        root = ((-b + math.Sqrt(delta)) / (2.0 * a))
        if root < 0 {
            return 0.0, vector3.Vector3{}, false
        }
    }

    root -= 0.0001
    p := ray.RayAt(root)
    normal := vector3.UnitVector(vector3.SubVector3(p, s.Center))

    return root, normal, true
}
