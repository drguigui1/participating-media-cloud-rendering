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
** 'Origin': Point shifted from the origin
*/
type VoxelGrid struct {
    VoxelSize float64
    NbVoxelX int
    NbVoxelY int
    NbVoxelZ int
    Origin vector3.Vector3

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
                   nbVoxelX,
                   nbVoxelY,
                   nbVoxelZ int,
                   origin vector3.Vector3) VoxelGrid {

    nbVoxels := nbVoxelX * nbVoxelY * nbVoxelZ
    voxels := make([]Voxel, nbVoxels)

    for i := 0; i < nbVoxels; i += 1 {
        voxels[i] = InitVoxel(0.0, 0.0, vector3.InitVector3(200.0 / 255.0,
                                                            100.0 / 255.0,
                                                            20.0 / 255.0))
    }

    return VoxelGrid{
        VoxelSize: voxelSize, // size of one voxel
        NbVoxelX: nbVoxelX,
        NbVoxelY: nbVoxelY,
        NbVoxelZ: nbVoxelZ,
        Origin: origin,
        Voxels: voxels,
    }
}
