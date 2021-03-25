package raymarch

import (
    //"fmt"
    "testing"
    "reflect"
    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/voxel_grid"
    
)

func TestRayMarchVoxelGrid(t *testing.T) {
    voxelGrid := voxel_grid.InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0))


    origin := vector3.InitVector3(2.5, 2.0, 0.0)
    dir := vector3.InitVector3(0.0, 0.0, -1.0)
    ray := ray.InitRay(origin, dir)

    
    
    points, hasHit := RayMarchVoxelGrid(ray, voxelGrid, 0.5)
    //fmt.Println(points)
    if !hasHit {
        t.Errorf("Error 'TestRayMarchVoxelGrid '")
        t.Errorf("res: %v\n", hasHit)
        t.Errorf("ref: %v\n", true)
    }
    m := []vector3.Vector3{
        vector3.InitVector3(2.5, 2, -1.0),
        vector3.InitVector3(2.5, 2, -1.5),
        vector3.InitVector3(2.5, 2, -2.0),
        vector3.InitVector3(2.5, 2, -2.5),
        vector3.InitVector3(2.5, 2, -3.0),
        vector3.InitVector3(2.5, 2, -3.5),
        vector3.InitVector3(2.5, 2, -4.0),
    }
    //fmt.Println(m)
    for i := 0; i < len(points); i+=1 {
        if !reflect.DeepEqual(points[i], m[i]) {
            t.Errorf("Error 'TestRayMarchVoxelGrid '")
            t.Errorf("res: %v\n", points[i])
            t.Errorf("ref: %v\n", m[i])
        }
    }
}
