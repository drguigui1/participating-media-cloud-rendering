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
*/
type VoxelGrid struct {
    VoxelSize float64
    nbVoxelX int
    nbVoxelY int
    nbVoxelZ int

    Voxels []Voxel
}

func InitVoxel(density, transmitivity float64, color vector3.Vector3) Voxel {
    return Voxel{
        Density: density,
        Transmitivity: transmitivity,
        Color: color,
    }
}

func InitVoxelGrid() VoxelGrid {
    // TODO
    return VoxelGrid{}
}
