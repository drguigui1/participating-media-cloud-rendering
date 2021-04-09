package background

import (
    "volumetric-cloud/ray"
    "volumetric-cloud/vector3"
)

func RenderGradient(ray ray.Ray) vector3.Vector3 {
    dir := vector3.UnitVector(ray.Direction);
    tmp := 0.5 * (dir.Y + 1.0);
    tmp2 := 1.0 - tmp
    return vector3.AddVector3(vector3.InitVector3(1.0 * tmp2, 1.0 * tmp2, 1.0 * tmp2), vector3.InitVector3(0.1 * tmp, 0.6 * tmp, 0.8 * tmp))
}
