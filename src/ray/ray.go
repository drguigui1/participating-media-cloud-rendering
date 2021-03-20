package ray

import (
    "volumetrical-cloud/vector3"
)

type Ray struct {
    Origin vector3.Vector3
    Direction vector3.Vector3
}

func InitRay(ori, dir vector3.Vector3) Ray {
    return Ray{
        Origin: ori,
        Direction: dir,
    }
}

func (r Ray) RayAt(t float64) vector3.Vector3 {
    return vector3.AddVector3(r.Origin, vector3.MulVector3(r.Direction, t))
}
