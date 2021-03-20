package ray

import (
    "testing"
    "reflect"

    "volumetrical-cloud/vector3"
)

func TestRayAt(t *testing.T) {
    origin := vector3.InitVector3(1, 1, 1)
    directory := vector3.InitVector3(-1, -1, -1)
    ray := InitRay(origin, directory)

    res := ray.RayAt(2)
    ref := vector3.InitVector3(-1, -1, -1)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error Test RayAt")
        t.Errorf("res: %v\n", res)
        t.Errorf("ref: %v\n", ref)
    }
}
