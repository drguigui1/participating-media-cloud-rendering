package voxel_grid

import (
    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/plane"
)

/*
** One voxel == one cube
** 'Density': density in the voxel
** 'Transmitivity': how the voxel transmit the light
** 'Color': color inside the voxel
*/
type Voxel struct {
    Density float64
    Transmitivity float64
    Color vector3.Vector3
}

/*
** 3D grid of voxels
** 'VoxelSize': size of one voxel (one cube)
** 'NbVoxelX': Number of lattice for the x axis: (NbVoxelsX - 1) = NbCubesX
** 'ShiftedPoint': Shift the voxel grid in the space
*/
type VoxelGrid struct {
    VoxelSize float64
    NbVerticeX int
    NbVerticeY int
    NbVerticeZ int

    Shift vector3.Vector3
    Voxels []Voxel
}

func InitVoxel(density, transmitivity float64, color vector3.Vector3) Voxel {
    return Voxel{
        Density: density,
        Transmitivity: transmitivity,
        Color: color,
    }
}

func InitVoxelGrid(voxelSize float64,
                   nbVerticeX,
                   nbVerticeY,
                   nbVerticeZ int,
                   shift vector3.Vector3) VoxelGrid {

    nbVertices := nbVerticeX * nbVerticeY * nbVerticeZ
    voxels := make([]Voxel, nbVertices)

    // Init voxels
    for i := 0; i < nbVertices; i += 1 {
        voxels[i] = InitVoxel(0.0, 0.0, vector3.InitVector3(200.0 / 255.0,
                                                            100.0 / 255.0,
                                                            20.0 / 255.0))
    }

    return VoxelGrid{
        VoxelSize: voxelSize, // size of one voxel
        NbVerticeX: nbVerticeX,
        NbVerticeY: nbVerticeY,
        NbVerticeZ: nbVerticeZ,
        Shift: shift,
        Voxels: voxels,
    }
}

func (vGrid VoxelGrid) ShiftToVoxelCoordinates(p vector3.Vector3) vector3.Vector3 {
    return vector3.SubVector3(p, vGrid.Shift)
}

func (vGrid VoxelGrid) ShiftToWorldCoordinates(voxelCoordinatePoint vector3.Vector3) vector3.Vector3 {
    return vector3.AddVector3(voxelCoordinatePoint, vGrid.Shift)
}

func (vGrid VoxelGrid) IsInsideVoxelGrid(p vector3.Vector3) bool {
    pVoxel := vGrid.ShiftToVoxelCoordinates(p)

    if pVoxel.X < 0 || pVoxel.X > (vGrid.VoxelSize * float64(vGrid.NbVerticeX - 1)) ||
       pVoxel.Y < 0 || pVoxel.Y > (vGrid.VoxelSize * float64(vGrid.NbVerticeY - 1)) ||
       pVoxel.Z < 0 || pVoxel.Z > (vGrid.VoxelSize * float64(vGrid.NbVerticeZ - 1)) {
        return false
    }

    return true
}

/*
** Get the first point that intersect the VoxelGrid
*/
func (vGrid VoxelGrid) IntersectFaces(ray ray.Ray) (float64, bool, vector3.Vector3) {
    edgeSizeX := vGrid.VoxelSize * float64(vGrid.NbVerticeX - 1)
    edgeSizeY := vGrid.VoxelSize * float64(vGrid.NbVerticeY - 1)
    edgeSizeZ := vGrid.VoxelSize * float64(vGrid.NbVerticeZ - 1)

    normals := []vector3.Vector3{
        vector3.InitVector3(0.0, 0.0, -1.0),
        vector3.InitVector3(0.0, 0.0, 1.0),
        vector3.InitVector3(-1.0, 0.0, 0.0),
        vector3.InitVector3(1.0, 0.0, 0.0),
        vector3.InitVector3(0.0, 1.0, 0.0),
        vector3.InitVector3(0.0, -1.0, 0.0),
    }

    points := [][]vector3.Vector3{
        { // perpendicular to z
            vGrid.Shift.Copy(),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y, vGrid.Shift.Z),
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z),
        },
        { // perpendicular to z
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y, vGrid.Shift.Z + edgeSizeZ),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y, vGrid.Shift.Z + edgeSizeZ),
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z + edgeSizeZ),
        },
        { // perpendicular to x
            vGrid.Shift.Copy(),
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z),
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y, vGrid.Shift.Z + edgeSizeZ),
        },
        { // perpendicular to x
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y, vGrid.Shift.Z),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y, vGrid.Shift.Z + edgeSizeZ),
        },
        { // perpendicular to y
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z),
            vector3.InitVector3(vGrid.Shift.X, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z + edgeSizeZ),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y + edgeSizeY, vGrid.Shift.Z + edgeSizeZ),
        },
        { // perpendicular to y
            vGrid.Shift.Copy(),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y, vGrid.Shift.Z),
            vector3.InitVector3(vGrid.Shift.X + edgeSizeX, vGrid.Shift.Y, vGrid.Shift.Z + edgeSizeZ),
        },
    }

    colors := []vector3.Vector3{
        vector3.InitVector3(200, 100, 10), // orange
        vector3.InitVector3(100, 200, 10), // green
        vector3.InitVector3(100, 100, 100), // grey
        vector3.InitVector3(10, 100, 100), // blue
        vector3.InitVector3(0, 0, 0), // black
        vector3.InitVector3(200, 11, 168), // purple
    }

    var t float64 = 0.0
    var hasHit bool = false
    var finalColor vector3.Vector3

    for i := 0; i < 6; i += 1 {
        // create plane object
        plane := plane.InitPlane(normals[i], colors[i], points[i][0], points[i][1], points[i][2])

        // intersect a face
        t, hasHit = plane.Hit(ray)

        if hasHit {
            finalColor = colors[i]
            break
        }
    }

    return t, hasHit, finalColor
}
