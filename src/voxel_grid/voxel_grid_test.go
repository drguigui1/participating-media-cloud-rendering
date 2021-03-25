package voxel_grid

import (
    "testing"
    "volumetric-cloud/vector3"
)

func TestInitVoxelGrid1(t *testing.T) {
    voxelGrid := InitVoxelGrid(1.0,
                  vector3.InitVector3(0.0, 0.0, -1.0),
                  vector3.InitVector3(5.0, 4.0, -4.0))

    refNbVerticesX := 6
    refNbVerticesY := 5
    refNbVerticesZ := 4

    if voxelGrid.NbVerticeX != refNbVerticesX {
        t.Errorf("Error 'TestInitVoxelGrid' nbVerticesX")
        t.Errorf("Res: %v\n", voxelGrid.NbVerticeX)
        t.Errorf("Ref: %v\n", refNbVerticesX)
    }

    if voxelGrid.NbVerticeY != refNbVerticesY {
        t.Errorf("Error 'TestInitVoxelGrid' nbVerticesY")
        t.Errorf("Res: %v\n", voxelGrid.NbVerticeY)
        t.Errorf("Ref: %v\n", refNbVerticesY)
    }

    if voxelGrid.NbVerticeZ != refNbVerticesZ {
        t.Errorf("Error 'TestInitVoxelGrid' nbVerticesZ")
        t.Errorf("Res: %v\n", voxelGrid.NbVerticeZ)
        t.Errorf("Ref: %v\n", refNbVerticesZ)
    }
}

func TestInitVoxelGrid2(t *testing.T) {
    voxelGrid := InitVoxelGrid(1.0,
                  vector3.InitVector3(0.0, 0.0, -1.0),
                  vector3.InitVector3(-5.0, 4.0, -4.0))

    refNbVerticesX := 6
    refNbVerticesY := 5
    refNbVerticesZ := 4

    if voxelGrid.NbVerticeX != refNbVerticesX {
        t.Errorf("Error 'TestInitVoxelGrid' nbVerticesX")
        t.Errorf("Res: %v\n", voxelGrid.NbVerticeX)
        t.Errorf("Ref: %v\n", refNbVerticesX)
    }

    if voxelGrid.NbVerticeY != refNbVerticesY {
        t.Errorf("Error 'TestInitVoxelGrid' nbVerticesY")
        t.Errorf("Res: %v\n", voxelGrid.NbVerticeY)
        t.Errorf("Ref: %v\n", refNbVerticesY)
    }

    if voxelGrid.NbVerticeZ != refNbVerticesZ {
        t.Errorf("Error 'TestInitVoxelGrid' nbVerticesZ")
        t.Errorf("Res: %v\n", voxelGrid.NbVerticeZ)
        t.Errorf("Ref: %v\n", refNbVerticesZ)
    }
}



func TestIsInsideVoxelGrid1(t *testing.T) {
    shift := vector3.InitVector3(0.0, 0.0, 0.0)
    oppositeCorner := vector3.InitVector3(3.0, 3.0, 3.0)

    newVoxelGrid := InitVoxelGrid(0.5, shift, oppositeCorner)

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
            t.Errorf("Error 'TestIsInsideVoxelGrid1'")
            t.Errorf("Res: %v\n", res)
            t.Errorf("Ref: %v\n", true)
        }
    }
}

func TestIsInsideVoxelGrid2(t *testing.T) {
    shift := vector3.InitVector3(1.0, 1.0, 1.0)
    oppositeCorner := vector3.InitVector3(3.0, 3.0, 3.0)

    newVoxelGrid := InitVoxelGrid(0.5, shift, oppositeCorner)

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
            t.Errorf("Error 'TestIsInsideVoxelGrid2' idx: %v", idx)
            t.Errorf("Res: %v\n", res)
            t.Errorf("Ref: %v\n", ref[idx])
        }
    }
}
