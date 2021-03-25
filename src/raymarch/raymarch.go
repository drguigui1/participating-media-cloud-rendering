package raymarch

import (
    "fmt"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
)

func RayMarchVoxelGrid(ray ray.Ray, voxelGrid voxel_grid.VoxelGrid, step float64) ([]vector3.Vector3, bool) {
    // Check if already inside
    var t float64
    var hasHit bool
    if voxelGrid.IsInsideVoxelGrid(ray.Origin) {
        fmt.Println("ray")
        t = 0
    } else {
        // Get first point (origin)
        t, hasHit, _ = voxelGrid.Hit(ray)

        if !hasHit {
            return []vector3.Vector3{}, false
        }

    }

    // allocate slice
    points := make([]vector3.Vector3, 0)

    // compute origin
    o := ray.RayAt(t + 0.001)
    fmt.Println(o)

    for voxelGrid.IsInsideVoxelGrid(o) {
        points = append(points, o)
        // TODO: add random value to step
        o = vector3.AddVector3(o, vector3.MulVector3(ray.Direction, step))
        fmt.Println(o)
    }

    return points, true
}
