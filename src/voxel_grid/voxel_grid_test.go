package voxel_grid

import (
    "testing"
    "volumetric-cloud/vector3"
)

func TestIsInsideVoxelGrid1(t *testing.T) {
    shift := vector3.InitVector3(0.0, 0.0, 0.0)

    newVoxelGrid := InitVoxelGrid(0.5, 5, 5, 5, shift)

    p1 := vector3.InitVector3(0.5, 0.5, 0.5)
    p2 := vector3.InitVector3(0.6, 0.5, 0.5)
    p3 := vector3.InitVector3(0.5, 1.5, 0.5)
    p4 := vector3.InitVector3(0.5, 1.5, 1.5)
    p5 := vector3.InitVector3(1.5, 1.5, 1.5)

    tests := []vector3.Vector3{
        p1, p2, p3, p4, p5,
    }

    for _, elm := range tests {
        res := newVoxelGrid.IsInsideVoxelGrid(elm)
        if res != true {
            t.Errorf("Error 'TestIsInsideVoxelGrid'")
            t.Errorf("Res: %v\n", res)
            t.Errorf("Ref: %v\n", true)
        }
    }
}

func TestIsInsideVoxelGrid2(t *testing.T) {
    shift := vector3.InitVector3(1.0, 1.0, 1.0)

    newVoxelGrid := InitVoxelGrid(0.5, 5, 5, 5, shift)

    p1 := vector3.InitVector3(0.5, 0.5, 0.5)
    p2 := vector3.InitVector3(0.6, 0.5, 2.5)
    p3 := vector3.InitVector3(0.5, 1.5, 0.5)
    p4 := vector3.InitVector3(1.5, 1.5, 1.5)
    p5 := vector3.InitVector3(2.5, 1.5, 2.9)

    tests := []vector3.Vector3{
        p1, p2, p3, p4, p5,
    }

    ref := []bool{
        false, false, false, true, true,
    }

    for idx, elm := range tests {
        res := newVoxelGrid.IsInsideVoxelGrid(elm)
        if res != ref[idx] {
            t.Errorf("Error 'TestIsInsideVoxelGrid'")
            t.Errorf("Res: %v\n", res)
            t.Errorf("Ref: %v\n", ref[idx])
        }
    }
}
