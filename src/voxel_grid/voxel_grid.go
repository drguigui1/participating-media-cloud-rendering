package voxel_grid

import (
    "volumetric-cloud/vector3"
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
func (vGrid VoxelGrid) IntersectFaces() float64 {
    // TODO
    return 0.0
}
