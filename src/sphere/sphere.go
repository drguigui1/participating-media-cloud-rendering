package sphere

import (
    "math"

    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/light"
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
    b := vector3.DotProduct(vector3.MulVector3Scalar(ray.Direction, 2.0), diff)
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

func (s Sphere) ComputeDiffuseGroundColor(lights []light.Light, groundColor, p vector3.Vector3, albedo float64) vector3.Vector3 {
    diffuseIntensity := vector3.InitVector3(0.0, 0.0, 0.0)
    for _, light := range lights {
        lightIntensity := light.Color

        lightDir := vector3.UnitVector(vector3.SubVector3(light.Position, p))
        normal := vector3.UnitVector(vector3.SubVector3(p, s.Center))
        lightImp := vector3.MulVector3Scalar(lightIntensity, math.Max(0.0, vector3.DotProduct(lightDir, normal)))
        diffuseIntensity.AddVector3(lightImp)
    }
    res := vector3.HadamarProduct(diffuseIntensity, groundColor)
    res.Mul(albedo)
    return res
}
