package voxel_grid

import (
    "testing"
    "reflect"
    "math"

    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
)

func TestInitVoxelGrid1(t *testing.T) {
    voxelGrid := InitVoxelGrid(1.0,
                  vector3.InitVector3(0.0, 0.0, -1.0),
                  vector3.InitVector3(5.0, 4.0, -4.0), 0.0)

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
                  vector3.InitVector3(-5.0, 4.0, -4.0), 0.0)

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

    newVoxelGrid := InitVoxelGrid(0.5, shift, oppositeCorner, 0.0)

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

    newVoxelGrid := InitVoxelGrid(0.5, shift, oppositeCorner, 0.0)

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

func TestHit1(t *testing.T) {
    o := vector3.InitVector3(0.5, 0.5, 0.0)
    d := vector3.InitVector3(0.0, 0.0, -1.0)
    ray := ray.InitRay(o, d)

    shift := vector3.InitVector3(0.0, 0.0, -2.0)
    oppositePoint := vector3.InitVector3(1.0, 1.0, -4.0)
    voxelGrid := InitVoxelGrid(0.5, shift, oppositePoint, 0.0)

    res, hasHit, _ := voxelGrid.Hit(ray)

    if math.Round(res) != 2.0 || !hasHit {
        t.Errorf("Error 'TestHit1'")
        t.Errorf("Res t: %v\n", res)
        t.Errorf("Ref t: %v\n", 2.0)
    }
}

func TestHit2(t *testing.T) {
    o := vector3.InitVector3(3.0, 0.5, -2.5)
    d := vector3.InitVector3(-1.0, 0.0, 0.0)
    ray := ray.InitRay(o, d)

    shift := vector3.InitVector3(0.0, 0.0, -2.0)
    oppositePoint := vector3.InitVector3(1.0, 1.0, -4.0)
    voxelGrid := InitVoxelGrid(0.5, shift, oppositePoint, 0.0)

    res, hasHit, _ := voxelGrid.Hit(ray)

    if math.Round(res) != 2.0 || !hasHit {
        t.Errorf("Error 'TestHit2'")
        t.Errorf("Res t: %v\n", res)
        t.Errorf("Ref t: %v\n", 2.0)
    }
}

func TestIsInsideVoxelGrid3(t *testing.T) {
    shift := vector3.InitVector3(0.0, 0.0, -4.0)
    oppositeCorner := vector3.InitVector3(5.0, 4.0, -1.0)

    newVoxelGrid := InitVoxelGrid(0.5, shift, oppositeCorner, 0.0)

    p1 := vector3.InitVector3(2.5, 2.0, 0.0)
    res := newVoxelGrid.IsInsideVoxelGrid(p1)
    if res != false {
        t.Errorf("Error 'TestIsInsideVoxelGrid2'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", false)
    }
}

func TestRayMarchVoxelGrid(t *testing.T) {
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0),
                               0.5)


    origin := vector3.InitVector3(2.5, 2.0, 0.0)
    dir := vector3.InitVector3(0.0, 0.0, -1.0)
    ray := ray.InitRay(origin, dir)

    points, hasHit := voxelGrid.RayMarch(ray)
    //fmt.Println(points)
    if !hasHit {
        t.Errorf("Error 'TestRayMarchVoxelGrid '")
        t.Errorf("res: %v\n", hasHit)
        t.Errorf("ref: %v\n", true)
    }
    m := []vector3.Vector3{
        vector3.InitVector3(2.5, 2, -1.001),
        vector3.InitVector3(2.5, 2, -1.501),
        vector3.InitVector3(2.5, 2, -2.001),
        vector3.InitVector3(2.5, 2, -2.501),
        vector3.InitVector3(2.5, 2, -3.001),
        vector3.InitVector3(2.5, 2, -3.501),
        vector3.InitVector3(2.5, 2, -4.001),
    }
    //fmt.Println(m)
    for i := 0; i < len(points); i+=1 {
        if Round3(points[i].X) != m[i].X ||
           Round3(points[i].Y) != m[i].Y ||
           Round3(points[i].Z) != m[i].Z {
            t.Errorf("Error 'TestRayMarchVoxelGrid'")
            t.Errorf("res: %v\n", points[i])
            t.Errorf("ref: %v\n", m[i])
        }
    }
}

func TestGetWorldPosition1(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetWorldPosition(vector3.InitVector3(0, 0, 0))

    if !reflect.DeepEqual(res, voxelGrid.Shift) {
        t.Errorf("Error 'TestGetWorldPosition1'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", voxelGrid.Shift)
    }
}

func TestGetWorldPosition2(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetWorldPosition(vector3.InitVector3(0, 0, 1))
    ref := vector3.InitVector3(voxelGrid.Shift.X, voxelGrid.Shift.Y, voxelGrid.Shift.Z + voxelGrid.VoxelSize)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestGetWorldPosition2'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetWorldPosition3(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetWorldPosition(vector3.InitVector3(3, 2, 1))
    ref := vector3.InitVector3(voxelGrid.Shift.X + 3 * voxelGrid.VoxelSize,
                               voxelGrid.Shift.Y + 2 * voxelGrid.VoxelSize,
                               voxelGrid.Shift.Z + voxelGrid.VoxelSize)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestGetWorldPosition3'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetVoxelIndex1(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetVoxelIndex(vector3.InitVector3(0.0, 0.0, -4.0))
    ref := vector3.InitVector3(0, 0, 0)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestGetVoxelIndex1'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetVoxelIndex2(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetVoxelIndex(vector3.InitVector3(1.0, 1.0, -3.0))
    ref := vector3.InitVector3(1, 1, 1)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestGetVoxelIndex2'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetVoxelIndex3(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetVoxelIndex(vector3.InitVector3(0.0, 1.0, -2.0))
    ref := vector3.InitVector3(0, 1, 2)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestGetVoxelIndex3'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetVoxelIndex4(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    res := voxelGrid.GetVoxelIndex(vector3.InitVector3(5.0, 4.0, -1.0))
    ref := vector3.InitVector3(5, 4, 3)

    if !reflect.DeepEqual(res, ref) {
        t.Errorf("Error 'TestGetVoxelIndex4'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetDensity1(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)
    voxelGrid.Voxels[0].Density = 0.2
    ref := 0.2
    res := voxelGrid.GetDensity(0,0,0)

    if ref != res {
        t.Errorf("Error 'TestGetDensity1'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestGetDensity2(t *testing.T) {
    // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)
    voxelGrid.Voxels[7].Density = 0.2
    ref := 0.2
    res := voxelGrid.GetDensity(1,1,0)

    if ref != res {
        t.Errorf("Error 'TestGetDensity1'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}

func TestSetTransmitivity1(t *testing.T) {
     // init the voxel grid for the test
    voxelGrid := InitVoxelGrid(1.0,
                               vector3.InitVector3(0.0, 0.0, -4.0),
                               vector3.InitVector3(5.0, 4.0, -1.0), 0.0)

    ref := 0.5
    voxelGrid.SetTransparency(0, 0, 1, ref)

    res := voxelGrid.Voxels[30].Transparency

    if ref != res {
        t.Errorf("Error 'TestSetTransparency1'")
        t.Errorf("Res: %v\n", res)
        t.Errorf("Ref: %v\n", ref)
    }
}
